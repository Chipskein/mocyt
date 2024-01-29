package main

import (
	"chipskein/yta-cli/internals/mpv"
	"chipskein/yta-cli/internals/ytdlp"
	"log"
)

func main() {
	var test = "https://www.youtube.com/watch?v=78ZLZfvzBNM"
	cmd, pipe, err := ytdlp.DownloadVideo(test)
	if err != nil {
		log.Fatalln(err)
	}
	err = mpv.Play(cmd, pipe)
	if err != nil {
		log.Fatalln(err)
	}
}
