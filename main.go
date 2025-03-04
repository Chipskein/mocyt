package main

import (
	"fmt"
	"os"

	"github.com/Chipskein/mocyt/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCMD := cobra.Command{
		Use:   "mocyt",
		Short: "Music from YT on console",
		Long:  `Simple terminal music player that streams audio from yt`,
	}
	cmd.LoginCmd.Flags().StringVarP(&cmd.CredentialsPath, "credentials", "c", cmd.CredentialsPath, "Path to credentials.json file Default is client_secret.json")
	cmd.LoginCmd.Flags().StringVarP(&cmd.TokenJsonPath, "token", "t", cmd.TokenJsonPath, "Path to token.json file Default is token.json")
	cmd.StartCmd.Flags().IntVarP(&cmd.SEARCH_MODE, "SEARCH_MODE", "s", cmd.SEARCH_MODE, "Sets SEARCH_MODE Default SCRAPPER_MODE  Set 2 to use API_MODE")
	cmd.StartCmd.Flags().StringVarP(&cmd.CredentialsPath, "credentials", "c", cmd.CredentialsPath, "Path to credentials.json file Default is client_secret.json")
	cmd.StartCmd.Flags().StringVarP(&cmd.TokenJsonPath, "token", "t", cmd.TokenJsonPath, "Path to token.json file Default is token.json")

	rootCMD.AddCommand(cmd.LoginCmd)
	rootCMD.AddCommand(cmd.StartCmd)
	rootCMD.AddCommand(cmd.KillCmd)
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
