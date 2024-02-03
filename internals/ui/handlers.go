package ui

import (
	"chipskein/yta-cli/internals/cache_handler"
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
		//check this what a brain fuck
		t.Current_player_info.Paused = !paused
		t.Current_player_info.Playing = paused
		//cache_handler.WriteInfo(t.Current_player_info)
	}
}

func handleVolumeDown(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume > 0 {
			mpv.SetVolume(current_volume - 1)
			t.Current_player_info.Volume = int32(current_volume - 1)
			//cache_handler.WriteInfo(t.Current_player_info)
			t.grid.Volumemixer.UpdatePercent(int(current_volume - 1))
		}
	}
}

func handleVolumeUp(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		current_volume, _ := mpv.GetVolume()
		if current_volume < 100 {
			mpv.SetVolume(current_volume + 1)
			t.Current_player_info.Volume = int32(current_volume + 1)
			//cache_handler.WriteInfo(t.Current_player_info)
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
		if t.Current_player_info.Playing {
			if t.grid.Progressbar.Root.Percent < 100 {
				secondsTotal := utils.ConvertStringToSeconds(t.Current_player_info.Duration)
				seconds := utils.ConvertStringToSeconds(t.Current_player_info.PlaybackTime)
				t.grid.Progressbar.Root.Percent = int(seconds * 100 / secondsTotal)
				t.Current_player_info.PercentProgresBar = int32(t.grid.Progressbar.Root.Percent)
			}
		}
	}
	t.UpdateScreen()
}
func handleEachSecond(t *TUI) {
	if mpv.CheckIfMpvIsRunning() {
		if t.Current_player_info.Playing {
			seconds, _ := mpv.GetPlayBackTimeSecond()
			time := utils.ConvertSecondsToString(int(seconds))
			t.Current_player_info.PlaybackTime = time
			t.grid.Plabackinfo.Update(time, t.Current_player_info.Duration)
		}
	}
	t.UpdateScreen()
}
func handleEach10Second(t *TUI) {
	cache_handler.WriteInfo(t.Current_player_info)
}
func HandleSelectedVideo(t *TUI, videoname string) {
	id, duration, videoname := utils.ParseListString(videoname)
	t.Current_player_info.ListString = videoname
	t.Current_player_info.Duration = duration
	t.Current_player_info.PlaybackTime = "0s"
	t.Current_player_info.PercentProgresBar = 0
	t.Current_player_info.Playing = true
	t.Current_player_info.Paused = false
	t.grid.Plabackinfo.Update(t.Current_player_info.PlaybackTime, duration)
	t.grid.Progressbar.Update(int(t.Current_player_info.PercentProgresBar), videoname, t.Current_player_info.Playing)
	go handlePlay(id)
	//gambers
	time.Sleep(1 * time.Second)
	mpv.SetVolume(float64(t.Current_player_info.Volume))
}
