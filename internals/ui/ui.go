package ui

import (
	"chipskein/yta-cli/internals/ui/components"
	"log"
	"time"

	tui "github.com/gizak/termui/v3"
)

type TUI struct {
	shouldRenderSearchBar bool
	grid                  *components.Grid
	uiEvents              <-chan tui.Event
	searchTxt             string
	tickerProgresBar      *<-chan time.Time
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
		}

	}
}

func (t *TUI) UpdateScreen() {
	tui.Render(t.grid.Root)
}

func StartUI() {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()
	var t = &TUI{}
	t.grid = components.Init()
	t.UpdateScreen()
	t.uiEvents = tui.PollEvents()
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.HandleTUIEvents()

}
