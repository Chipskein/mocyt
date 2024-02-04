package cmd

import (
	"chipskein/yta-cli/internals/repositories"
	"chipskein/yta-cli/internals/repositories/youtube"
	ytsearchscrapper "chipskein/yta-cli/internals/repositories/yt-search-scrapper"
	"chipskein/yta-cli/internals/ui"
	"context"

	"github.com/spf13/cobra"
)

const SCRAPPER_MODE = 1
const API_MODE = 2

var SEARCH_MODE int = SCRAPPER_MODE

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Music from YT on console",
	Long:  `Simple terminal music player that streams audio from yt`,
	Run: func(cmd *cobra.Command, args []string) {
		var repo repositories.YoutubeRepository
		switch SEARCH_MODE {
		case SCRAPPER_MODE:
			scrapper_repo := &ytsearchscrapper.YoutubeScrapper{}
			repo = scrapper_repo
		case API_MODE:
			ctx := context.Background()
			api_repo, err := youtube.Init(ctx, CredentialsPath, TokenJsonPath)
			if err != nil {
				panic(err)
			}
			repo = api_repo
		}
		ui.StartUI(repo)
	},
}
