package ui

import (
	"chipskein/yta-cli/internals/mpv"
	"chipskein/yta-cli/internals/utils"
	"chipskein/yta-cli/internals/ytdlp"
	"fmt"
	"time"
)

func handlePause(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		paused, _ := mpv.CheckMpvPaused()
		mpv.Pause(!paused)
		t.grid.Progressbar.Update(t.grid.Progressbar.Root.Percent, t.grid.Progressbar.Root.Label, paused)
	}
}

func handleVolumeDown(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume > 0 {
			mpv.SetVolume(current_volume - 1)
			t.grid.Volumemixer.UpdatePercent(int(current_volume - 1))
		}
	}
}

func handleVolumeUp(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume < 100 {
			mpv.SetVolume(current_volume + 1)
			t.grid.Volumemixer.UpdatePercent(int(current_volume + 1))
		}
	}
}
func handleMute(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		muted, _ := mpv.CheckMpvMute()
		mpv.Mute(!muted)
		t.grid.Volumemixer.SetMute(!muted)
	}
}

func handlePlay(videoID string) {
	if mpv.CheckIfMpvIsRunning() {
		mpv.Stop()
	}
	var yt_url = fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)
	cmd, stout, _ := ytdlp.DownloadVideo(yt_url)
	mpv.Play(cmd, stout)
}

func handleProgressBar(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		isPlaying, _ := mpv.CheckIfIsPlaying()
		if isPlaying {
			if t.grid.Progressbar.Root.Percent < 100 {
				t.grid.Progressbar.Root.Percent++
			}
		}
	}
	t.UpdateScreen()
}
func handleEachSecond(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		isPlaying, _ := mpv.CheckIfIsPlaying()
		if isPlaying {
			seconds, _ := mpv.GetPlayBackTimeSecond()
			time := utils.ConvertSecondsToString(int(seconds))
			t.grid.Plabackinfo.Update(time, t.durationString)
		}
	}
	t.UpdateScreen()
}
func HandleSelectedVideo(t *TUI, videoname string) {
	id, videoname, _, duration, err := utils.ParseListString(videoname)
	if err != nil {
		return
	}
	t.durationString = duration
	t.grid.Plabackinfo.Update("0s", duration)
	t.grid.Progressbar.Update(0, videoname, true)
	t.tickerProgresBar = &time.NewTicker(15 * time.Second).C
	go handlePlay(id)
}
