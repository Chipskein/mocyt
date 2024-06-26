## MOCYT(Music on Console from YT)
### Description
Simple terminal player that plays music from yt 

https://github.com/Chipskein/mocyt/assets/47486707/f43cc476-9bdd-4f29-9a6a-4259ed821878

### Requirements
 * Go
 * GNU Make
 * [yt-dlp](https://github.com/yt-dlp/yt-dlp)
 * [mpv](https://github.com/mpv-player/mpv)
### Install
**Linux:**
  After installing the required programs, run:

      git clone https://github.com/Chipskein/mocyt.git
      cd mocyt
      sudo make install
### Usage
     mocyt [command]
### Commands
| Command   |      Supported Flags      |  Description |
|------------|:---------------------------:|------:|
| start   | -c -t -s |  Will Start or Resume mocyt |
| login   | -c -t    |  Will create token.json for API_MODE |
| kill    | None     | Will kill MPV currently Playing and will clean up cached playback information |
### Flags
|      Flag     |  Shorhand |           Default            | Description |
|---------------|:---------:|-----------------------------:|------------:|
| --credentials | -c        | "./client_credentials.json"  | Path to Google Oauth2 client credentials to use YoutubeDataAPI |
| --token       | -t        | "./token.json"               | Path to Google Oauth2 token generated by mocyt login command   |
| --SEARCH_MODE | -s        | 1                            | Sets SEARCH_MODE flag |

### SEARCH_MODE
MOCYT supports two modes for performing YouTube searches:
 * SCRAPPER_MODE(1):
     * Utilizes a web scraper to get results from YouTube search.
 * API_MODE(2):
     * Utilizes YouTube Data API to get results from YouTube search, requiring additional configuration to utilize.
### Setup to use API_MODE
 * Create Google OAuth Client Credentials at Google Console [you can follow this tutorial from google](https://support.google.com/cloud/answer/6158849)
 * Run login command to create token.json
   
       mocyt start -c <PATH_TO_CLIENT_CREDENTIALS> -t <PATH_TO_TOKEN_JSON>
 * Run start
   
        mocyt start -s 2 -c <PATH_TO_CLIENT_CREDENTIALS> -t <PATH_TO_TOKEN_JSON>
 
### TODO
  * [ ] Containerize Application with docker
  * [ ] Deploy Image to Docker.hub
