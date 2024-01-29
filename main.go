package main

import (
	"chipskein/yta-cli/internals/mpv"
	"chipskein/yta-cli/internals/ytdlp"
	"log"
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
	go func() {
		//time.Sleep(time.Second * 10)
		/*
			err := mpv.Stop()
			if err != nil {
				log.Fatalln(err)
			}

			err = mpv.Pause(true)
			if err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Second * 10)
			err = mpv.Pause(false)
			if err != nil {
				log.Fatalln(err)
			}
		*/
	}()
	for {

	}
}
