package cmd

import (
	"chipskein/yta-cli/internals"

	"github.com/spf13/cobra"
)

var KillCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill mfytoc",
	Long:  `Will kill any instance of MPV currently playing and will delete cached files and mpv socket`,
	Run: func(cmd *cobra.Command, args []string) {
		internals.KillThemAll()
	},
}
