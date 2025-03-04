package ui

import (
	"log"
	"time"

	"github.com/Chipskein/mocyt/internals/cache_handler"
	"github.com/Chipskein/mocyt/internals/repositories"
	"github.com/Chipskein/mocyt/internals/ui/components"

	tui "github.com/gizak/termui/v3"
)

type TUI struct {
	repository            repositories.YoutubeRepository
	shouldRenderSearchBar bool
	grid                  *components.Grid
	uiEvents              <-chan tui.Event
	Current_player_info   *cache_handler.PlayerInformation
	tickerProgresBar      *<-chan time.Time
	tickerSecond          *<-chan time.Time
}

func (t *TUI) HandleTUIEvents() {
	for {
		select {
		case e := <-t.uiEvents:
			if !t.shouldRenderSearchBar {
				exit := HandleUserCommands(t, e)
				if exit {
					return
				}
			} else {
				HandleSearchInputEvents(t, e)
			}

		case <-*t.tickerProgresBar:
			handleProgressBar(t)
		case <-*t.tickerSecond:
			handleEachSecond(t)
		}

	}
}

func (t *TUI) UpdateScreen() {
	tui.Render(t.grid.Root)
}

func StartUI(repository repositories.YoutubeRepository) {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()
	var t = &TUI{repository: repository}
	t.grid = components.Init()
	cache := cache_handler.CheckIfCacheFileExists()
	if cache {
		cached_info := cache_handler.ReadInfo()
		t.Current_player_info = cached_info
		videos := []string{t.Current_player_info.SearchTxt}
		videos = append(videos, t.Current_player_info.SearchResults...)
		t.grid.Videolist.Update(videos, "Welcome!")
		t.grid.Plabackinfo.Update(t.Current_player_info.PlaybackTime, t.Current_player_info.Duration)
		t.grid.Volumemixer.UpdatePercent(int(t.Current_player_info.Volume))
		if t.Current_player_info.Muted {
			t.grid.Volumemixer.SetMute(t.Current_player_info.Muted)
		}
		t.grid.Progressbar.Update(int(t.Current_player_info.PercentProgresBar), t.Current_player_info.ListString, t.Current_player_info.Playing)
	} else {
		t.Current_player_info = &cache_handler.PlayerInformation{
			ListString:        "",
			PlaybackTime:      "0s",
			Duration:          "0s",
			PercentProgresBar: 0,
			Volume:            100,
			Paused:            false,
			Playing:           false,
			SearchTxt:         "",
			Muted:             false,
			SearchResults:     []string{},
		}
	}
	t.UpdateScreen()
	t.uiEvents = tui.PollEvents()
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.tickerSecond = &time.NewTicker(time.Second).C
	t.HandleTUIEvents()

}
