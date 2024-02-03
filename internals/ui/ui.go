package ui

import (
	"chipskein/yta-cli/internals/cache_handler"
	"chipskein/yta-cli/internals/repositories/youtube"
	"chipskein/yta-cli/internals/ui/components"
	"log"
	"time"

	tui "github.com/gizak/termui/v3"
)

type TUI struct {
	repository            *youtube.YoutubeRepository
	shouldRenderSearchBar bool
	grid                  *components.Grid
	uiEvents              <-chan tui.Event
	searchTxt             string
	Current_player_info   *cache_handler.PlayerInformation
	tickerProgresBar      *<-chan time.Time
	tickerSecond          *<-chan time.Time
	ticker10Second        *<-chan time.Time
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
		case <-*t.ticker10Second:
			handleEach10Second(t)
		}

	}
}

func (t *TUI) UpdateScreen() {
	tui.Render(t.grid.Root)
}

func StartUI(repository *youtube.YoutubeRepository) {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()
	var t = &TUI{repository: repository}
	t.grid = components.Init()
	//check cached info
	cache := cache_handler.CheckIfCacheFileExists()
	if cache {
		cached_info := cache_handler.ReadInfo()
		t.Current_player_info = cached_info
		t.grid.Plabackinfo.Update(t.Current_player_info.PlaybackTime, t.Current_player_info.Duration)
		t.grid.Volumemixer.UpdatePercent(int(t.Current_player_info.Volume))
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
			PidMPV:            ""}
	}
	t.UpdateScreen()
	t.uiEvents = tui.PollEvents()
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.tickerSecond = &time.NewTicker(time.Second).C
	t.ticker10Second = &time.NewTicker(10 * time.Second).C
	t.HandleTUIEvents()

}
