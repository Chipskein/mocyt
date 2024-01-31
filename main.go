package main

import (
	"chipskein/yta-cli/internals/repositories"
	"chipskein/yta-cli/internals/ui"
	"context"
)

func main() {
	ctx := context.Background()
	repo, err := repositories.Init(ctx, "client_secret.json")
	if err != nil {
		panic(err)
	}
	ui.StartUI(repo)
}
