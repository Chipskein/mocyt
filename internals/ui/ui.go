package ui

import (
	"chipskein/yta-cli/internals/cache_handler"
	"chipskein/yta-cli/internals/repositories"
	"chipskein/yta-cli/internals/ui/components"
	"log"
	"time"

	tui "github.com/gizak/termui/v3"
)

type TUI struct {
	repository            *repositories.YoutubeRepository
	shouldRenderSearchBar bool
	grid                  *components.Grid
	uiEvents              <-chan tui.Event
	searchTxt             string
	current_player_info   *cache_handler.PlayerInformation
	durationString        string //remove from here after
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

func StartUI(repository *repositories.YoutubeRepository) {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()
	var t = &TUI{repository: repository}
	//check cached info
	cache := cache_handler.CheckIfCacheFileExists()
	if cache {
		cached_info := cache_handler.ReadInfo()
		t.current_player_info = cached_info
	} else {
		t.current_player_info = &cache_handler.PlayerInformation{
			ListString:        "",
			PlaybackTime:      "0s",
			Duration:          "0s",
			PercentProgresBar: 0,
			Volume:            0,
			Paused:            false,
			Playing:           false,
			PidMPV:            ""}
	}
	t.grid = components.Init()
	t.UpdateScreen()
	t.uiEvents = tui.PollEvents()
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.tickerSecond = &time.NewTicker(time.Second).C
	t.durationString = "0s"
	t.HandleTUIEvents()

}
