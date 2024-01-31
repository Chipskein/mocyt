package repositories

import (
	"chipskein/yta-cli/internals/utils"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type YoutubeRepository struct {
	Service *youtube.Service
}

type Video struct {
	Id          string
	Title       string
	Duration    string
	ChannelName string
}

func (ytr *YoutubeRepository) ListVideos(searchTxt string) ([]string, error) {
	var videos []string
	service := ytr.Service
	if service == nil {
		return videos, errors.New("nil pointer at services in ytr")
	}
	call := service.Search.List([]string{"snippet"}).Q(searchTxt).Type("video")
	response, err := call.Do()
	if err != nil {
		return videos, err
	}
	for _, video := range response.Items {
		var id = video.Id.VideoId
		call := service.Videos.List([]string{"snippet", "contentDetails"}).Id(id)
		response, err := call.Do()
		if err != nil {
			return videos, err
		}
		if len(response.Items) == 0 {
			return videos, fmt.Errorf("could not found any video with id:%s", id)
		}
		var title = response.Items[0].Snippet.Title
		var channelName = response.Items[0].Snippet.ChannelTitle
		var duration = utils.ConvertPTISO8061(response.Items[0].ContentDetails.Duration)
		var video = Video{
			Id:          id,
			Title:       title,
			Duration:    duration,
			ChannelName: channelName}
		videos = append(videos, fmt.Sprintf("%s\nID:%s\nDuration:%s\nChannel:%s\n", video.Title, video.Id, video.Duration, video.ChannelName))

	}
	return videos, nil
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("youtube-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Init(ctx context.Context, credentials_path string) (*YoutubeRepository, error) {
	var result = &YoutubeRepository{}
	//"client_secret.json"
	b, err := ioutil.ReadFile(credentials_path)
	if err != nil {
		return result, err
	}
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		return result, err
	}
	client := getClient(ctx, config)
	_, err = os.Stat("token.json")
	if err == os.ErrNotExist {
		log.Println("Token file does not exist, creating one")
		token := getTokenFromWeb(config)
		saveToken("token.json", token)
	}

	service, err := youtube.New(client)
	result.Service = service
	return result, nil
}
