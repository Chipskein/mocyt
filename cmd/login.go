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
	Short: "Login with youtube Account",
	Long: `You need to login with youtube to utilize this program because it uses Youtube Data API 
	for search videos`,
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
