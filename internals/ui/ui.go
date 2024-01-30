package ui

import (
	"log"
	"sync"
	"time"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var wg sync.WaitGroup

type TUI struct {
	grid             *tui.Grid
	ticker           *<-chan time.Time
	tickerProgresBar *<-chan time.Time
	uiEvents         <-chan tui.Event
	progressBar      *widgets.Gauge
	volumeBar        *widgets.Gauge
	videolist        *widgets.List
	p                *widgets.Paragraph
}

func (t *TUI) RenderFileList() {
	videolist := widgets.NewList()
	videolist.Rows = []string{
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)"}
	videolist.Title = "persona 5 whims of fate"
	videolist.TitleStyle.Fg = tui.ColorWhite
	videolist.SelectedRowStyle.Fg = tui.ColorBlack
	videolist.SelectedRowStyle.Bg = tui.ColorWhite
	videolist.TextStyle.Fg = tui.ColorWhite
	t.videolist = videolist
}

func (t *TUI) RenderProgressBar() {

	processBar := widgets.NewGauge()
	processBar.TitleStyle.Fg = tui.ColorWhite
	processBar.Percent = 0
	processBar.Label = " "
	processBar.BarColor = tui.ColorWhite
	processBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	t.progressBar = processBar
}

func (t *TUI) SetupGrid() {
	grid := tui.NewGrid()
	termWidth, termHeight := tui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		tui.NewRow(1.6/2,
			tui.NewCol(1.5/2, t.videolist)),
		tui.NewRow(0.2/2,
			tui.NewCol(1.5/2, t.p),
			tui.NewCol(0.5/2, t.volumeBar)),
		tui.NewRow(0.18/2, t.progressBar),
	)
	t.grid = grid
	t.RenderUI()
}
func (t *TUI) HandleTUIEvents() {
	for {
		select {
		case e := <-t.uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "<Enter>":
				if !t.isSearching {
					t.HandleSelectedFile(t.videolist.Rows[t.videolist.SelectedRow])
					t.RenderUI()
				}
			case "<Down>", "j":
				t.videolist.ScrollDown()
			case "<Up>", "k":
				t.videolist.ScrollUp()
			case "<End>":
				t.videolist.ScrollBottom()
			case "<Home>":
				t.videolist.ScrollTop()
			case "<PageDown>":
				t.videolist.ScrollHalfPageDown()
			case "<PageUp>":
				t.videolist.ScrollHalfPageUp()
			case "<Space>": //pause
			case ",": //volume up
			case ".": //volume down
			case "m": //mute
			case "<Resize>":
				payload := e.Payload.(tui.Resize)
				t.grid.SetRect(0, 0, payload.Width, payload.Height)
				tui.Clear()
				t.RenderUI()
			}

		case <-*t.ticker:
			t.RenderUI()
		case <-*t.tickerProgresBar:
			t.RenderUI()
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
	t.progressBar.Percent = 0
	t.progressBar.Title = "|> Playing"
	t.progressBar.Label = videoname
	//var now = t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second)
	//var duration = t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second)
	//t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", now, duration)
	//t.tickerProgresBar = &time.NewTicker(duration / 100).C

}
func (t *TUI) RenderUI() {
	tui.Render(t.grid)
}

func StartUI(CURRENT_DIRECTORY string, DEFAULT_DIRECTORY string, ShowHiddenFiles bool) {
	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer tui.Close()

	var t = &TUI{}
	//t.repo = &repositories.LocalRepository{CURRENT_DIRECTORY: CURRENT_DIRECTORY, DEFAULT_DIRECTORY: DEFAULT_DIRECTORY, ShowHiddenFiles: ShowHiddenFiles}

	go t.RenderFileList()
	go t.RenderProgressBar()
	wg.Add(5)
	wg.Done()
	time.Sleep(time.Millisecond * 500)
	t.SetupGrid()
	t.uiEvents = tui.PollEvents()
	t.ticker = &time.NewTicker(time.Microsecond).C
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.HandleTUIEvents()

}
