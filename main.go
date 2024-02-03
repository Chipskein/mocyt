package main

import (
	"chipskein/yta-cli/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCMD := cobra.Command{
		Use:   "mfytoc",
		Short: "Music from YT on console",
		Long:  `Simple terminal music player that streams audio from yt`,
	}
	rootCMD.AddCommand(cmd.LoginCmd)
	rootCMD.AddCommand(cmd.StartCmd)
	rootCMD.AddCommand(cmd.KillCmd)
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
