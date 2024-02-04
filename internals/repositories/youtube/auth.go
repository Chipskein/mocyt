package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

var CONFIG = &oauth2.Config{}
var KillChannel = make(chan bool)
var token = &oauth2.Token{}

func GenAuthLink() {
	authURL := CONFIG.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+"authorization code: \n%v\n", authURL)
}
func ExchangeCode(code string) (token *oauth2.Token, err error) {
	tok, err := CONFIG.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return tok, nil
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
func Login(ctx context.Context, credentials_path string, token_json_path string) (err error) {
	//"client_secret.json"
	//https: //stackoverflow.com/questions/27585412/can-i-really-not-ship-open-source-with-client-id
	//fuck youtube
	b, err := os.ReadFile(credentials_path)
	if err != nil {
		return err
	}
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		return err
	}
	//fuck google console
	config.RedirectURL = "http://localhost:5000"
	CONFIG = config
	go GenAuthLink()
	time.Sleep(1 * time.Second)
	go InitServer(token_json_path)
	for {
		select {
		case <-KillChannel:
			client := CONFIG.Client(ctx, token)
			_, err := youtube.New(client)
			if err != nil {
				return err
			}
			return nil
		}

	}
}
func Init(ctx context.Context, credentials_path string, TokenJsonPath string) (*YoutubeDataAPI, error) {
	b, err := os.ReadFile(credentials_path)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(TokenJsonPath)
	if os.IsNotExist(err) {
		fmt.Println("Token file does not exist, please use login command do generate one")
		os.Exit(0)
	} else if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		return nil, err
	}
	token, err := tokenFromFile(TokenJsonPath)
	if err != nil {
		return nil, err
	}
	client := config.Client(ctx, token)
	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}
	return &YoutubeDataAPI{Service: service}, nil
}
