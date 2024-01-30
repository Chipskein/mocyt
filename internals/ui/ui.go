package ui

import (
	"chipskein/yta-cli/internals/ui/components"
	"log"
	"sync"

	tui "github.com/gizak/termui/v3"
)

var wg sync.WaitGroup

type TUI struct {
	shouldRenderSearchBar bool
	grid                  *components.Grid
	uiEvents              <-chan tui.Event
	searchTxt             string
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
		}

	}
}
func (t *TUI) HandleSelectedFile(videoname string) {
	//extract VideoID from filename
	//var file = t.repo.Files[filename]
	//if file.IsADirectory {
	//	t.repo.CURRENT_DIRECTORY = file.FullPath
	//	t.filelist.Rows = t.repo.ListFiles()
	//		t.filelist.Title = t.repo.CURRENT_DIRECTORY
	//		return
	//}
	//if t.player != nil {
	//	player.OLD_PERCENT = t.volumeBar.Percent
	//	player.OLD_VOLUME = t.player.Volume.Volume
	//	persistInfo = true
	//	t.player.Stop()
	//}
	//var f = repositories.ReadFile(file.FullPath)
	//streamer, format, _ := decoder.Decode(f, file.Extension)
	//t.player = player.InitPlayer(format.SampleRate, streamer, f)
	//go t.player.Play()
	//t.progressBar.Percent = 0
	//t.progressBar.Title = "|> Playing"
	//t.progressBar.Label = videoname
	//var now = t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second)
	//var duration = t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second)
	//t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", now, duration)
	//t.tickerProgresBar = &time.NewTicker(duration / 100).C

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
	t.HandleTUIEvents()
}
