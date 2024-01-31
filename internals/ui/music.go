package ui

import (
	"chipskein/yta-cli/internals/mpv"
	"chipskein/yta-cli/internals/utils"
	"chipskein/yta-cli/internals/ytdlp"
	"fmt"
)

func handlePause(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		paused, _ := mpv.CheckMpvPaused()
		mpv.Pause(!paused)
		t.UpdateScreen()
	}
}

// func handleStop()   {}
func handleVolumeDown(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume > 0 {
			mpv.SetVolume(current_volume - 1)
		}
		t.UpdateScreen()
	}
}

func handleVolumeUp(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume < 100 {
			mpv.SetVolume(current_volume + 1)
		}
		t.UpdateScreen()
	}
}

// func handleMute()   {}
func handlePlay(videoID string) {
	if mpv.CheckIfMpvIsRunning() {
		mpv.Stop()
	}
	var yt_url = fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	cmd, stout, _ := ytdlp.DownloadVideo(yt_url)
	mpv.Play(cmd, stout)

}

func HandleSelectedVideo(t *TUI, videoname string) {
	id, videoname, _, duration, err := utils.ParseListString(videoname)
	if err != nil {
		return
	}
	go handlePlay(id)
	t.grid.Plabackinfo.Update(0, duration)
	t.grid.Progressbar.Update(0, videoname)
	t.UpdateScreen()
}
