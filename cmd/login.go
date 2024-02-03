package cmd

import (
	"chipskein/yta-cli/internals/repositories/youtube"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with youtube Account",
	Long: `You need to login with youtube to utilize this program because it uses Youtube Data API 
	for search videos`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		err := youtube.Login(ctx, "client_secret.json")
		if err != nil {
			fmt.Println(err)
			//panic(err)
		}
		fmt.Println("Success")
	},
}
