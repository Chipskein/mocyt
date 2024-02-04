package cmd

import (
	ytsearchscrapper "chipskein/yta-cli/internals/repositories/yt-search-scrapper"
	"chipskein/yta-cli/internals/ui"

	"github.com/spf13/cobra"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Music from YT on console",
	Long:  `Simple terminal music player that streams audio from yt`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			ctx := context.Background()
			repo, err := youtube.Init(ctx, CredentialsPath, TokenJsonPath)
			if err != nil {
				panic(err)
			}*/
		repo := &ytsearchscrapper.YoutubeScrapper{}
		ui.StartUI(repo)
	},
}
