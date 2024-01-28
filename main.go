// Sample Go code for user authorization

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"google.golang.org/api/youtube/v3"
	//"google.golang.org/api/youtube/v3"
)

const missingClientSecretsMessage = `Please configure OAuth 2.0`

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

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func channelsListByUsername(service *youtube.Service, part string, forUsername string) {
	parr_list := []string{part}
	call := service.Channels.List(parr_list)
	call = call.ForUsername(forUsername)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Statistics.ViewCount))
}

type Video struct {
	Id          string
	Title       string
	Duration    int
	ChannelId   string
	ChannelName string
	URL         string
}

func videosListByName(service *youtube.Service, searchTxt string) {
	parr_list := []string{"snippet"}
	//var orderBy = "viewCount"
	var searchType = "video"
	call := service.Search.List(parr_list).Q(searchTxt).Type(searchType) //.Order(orderBy)
	response, err := call.Do()
	handleError(err, "")
	for _, video := range response.Items {
		var id = video.Id.VideoId
		call := service.Videos.List([]string{"snippet", "contentDetails"}).Id(id)
		response, err := call.Do()
		handleError(err, "")
		var title = response.Items[0].Snippet.Title
		var channelId = response.Items[0].Snippet.ChannelId
		var channelName = response.Items[0].Snippet.ChannelTitle
		var thumbnail = response.Items[0].Snippet.Thumbnails.Default.Url
		var duration = response.Items[0].ContentDetails.Duration
		var URL = fmt.Sprintf("https://www.youtube.com/watch?v=%s", id)
		fmt.Printf("%s\n %s\n %s\n %s\n %s\n %s\n %s\n _______\n", id, title, thumbnail, duration, channelId, channelName, URL)
		//fmt.Printf("%+v\n", response.Items[0].Snippet)
	}
}
func main() {
	/*
		ctx := context.Background()

		b, err := ioutil.ReadFile("client_secret.json")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved credentials
		// at ~/.credentials/youtube-go-quickstart.json
		config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		client := getClient(ctx, config)
		service, err := youtube.New(client)

		handleError(err, "Error creating YouTube client")
		//http://localhost/?state=state-token&code=4/0AfJohXlnVNVlvHpDexdL4poPI0_edIsOo5nkxJVnLXEVrJzANx3ueeetVDh7-MmzFOrl4w&scope=https://www.googleapis.com/auth/youtube.readonly
		//channelsListByUsername(service, "snippet,contentDetails,statistics", "GoogleDevelopers")
		videosListByName(service, "Persona 5 whims of fate")
	*/
	/*
		TODO:
		* Melhorar Sistema de login
			* Inicia um servidor que vai receber o endpoint do auth para conseguir pegar o token
			* Matar servidor sem matar o processo
			* Fazer cache do token de acesso
		* Fazer UI baseado no mocg
		Conferir JSON IPC do mpv
		https://github.com/mpv-player/mpv/blob/master/DOCS/man/ipc.rst
	*/
	//	var yt_test = ""
	//	_ := fmt.Sprintf("youtube-dl -f bestaudio 'https://www.youtube.com/watch?v=LRK6hjBZfLs' -o - 2>/dev/null | ffplay -nodisp -autoexit -i - &>/dev/null", yt_test)
	//fmt.Println(url)
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	path, err := exec.LookPath("youtube-dl")
	if err != nil {
		log.Fatal("LookPath: ", err)
	}
	fmt.Println(path)
	var commandArgs = []string{
		"-f",
		"bestaudio",
		"https://www.youtube.com/watch?v=LRK6hjBZfLs",
		"-o",
		"-"}

	cmd := exec.Command(path, commandArgs...)
	cmd.Stderr = os.Stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Pipe: ", err)
	}

	path, err = exec.LookPath("mpv")
	if err != nil {
		log.Fatal("LookPath: ", err)
	}
	fmt.Println(path)
	cmdArguments := []string{
		"-",
		"-input-ipc-server=/tmp/mpv-socket"}
	cmd2 := exec.Command(path, cmdArguments...)
	cmd2.Stdin = pipe
	cmd2.Stdout = devnull
	cmd2.Stderr = os.Stderr
	fmt.Println(cmd.String())
	fmt.Println(cmd2.String())

	cmd.Start()
	cmd2.Start()
	cmd.Wait()
	cmd2.Wait()

}
