package cmd

import (
	"chipskein/yta-cli/internals/repositories/youtube"
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var CredentialsPath string = "client_secret.json"
var TokenJsonPath string = "token.json"
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Create a token.json to utilize API_MODE",
	Long:  `To use API_MODE you need to create a token.json file to access YoutubeDataAPI`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := youtube.Login(ctx, CredentialsPath, TokenJsonPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		fmt.Println("Success")
	},
}
