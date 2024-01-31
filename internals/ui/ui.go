package ui

import (
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
	t.grid = components.Init()
	t.UpdateScreen()
	t.uiEvents = tui.PollEvents()
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.tickerSecond = &time.NewTicker(time.Second).C
	t.durationString = "0s"
	t.HandleTUIEvents()

}
