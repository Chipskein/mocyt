package main

import (
	"chipskein/yta-cli/internals/mpv"
	"chipskein/yta-cli/internals/ytdlp"
	"log"
	"time"
)

func main() {
	go func() {
		var test = "https://www.youtube.com/watch?v=78ZLZfvzBNM"
		cmd, pipe, err := ytdlp.DownloadVideo(test)
		if err != nil {
			log.Fatalln(err)
		}
		err = mpv.Play(cmd, pipe)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	for {
		time.Sleep(time.Second * 10)
		//testar outras coisas
	}
}
