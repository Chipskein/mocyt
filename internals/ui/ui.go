package ui

import (
	"fmt"
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
	queuelist        *widgets.List
	p                *widgets.Paragraph
	//repo             *repositories.LocalRepository
	//player           *player.PlayerController
}

func (t *TUI) RenderFileList() {
	videolist := widgets.NewList()
	videolist.Rows = []string{
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)"}
	videolist.Title = "persona 5 whims of fate" //t.repo.CURRENT_DIRECTORY
	videolist.TitleStyle.Fg = tui.ColorWhite
	videolist.SelectedRowStyle.Fg = tui.ColorBlack
	videolist.SelectedRowStyle.Bg = tui.ColorWhite
	videolist.TextStyle.Fg = tui.ColorWhite
	t.videolist = videolist
}
func (t *TUI) RenderQueueList() {
	queuelist := widgets.NewList()
	queuelist.Rows = []string{
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)",
		"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate (Bfr's OST)"}
	queuelist.Title = "Playback Queue" //t.repo.CURRENT_DIRECTORY
	queuelist.TitleStyle.Fg = tui.ColorWhite
	queuelist.SelectedRowStyle.Fg = tui.ColorBlack
	queuelist.SelectedRowStyle.Bg = tui.ColorWhite
	queuelist.TextStyle.Fg = tui.ColorWhite
	t.queuelist = queuelist
}
func (t *TUI) RenderVolumeMixer() {
	volumeBar := widgets.NewGauge()
	volumeBar.TitleStyle.Fg = tui.ColorWhite
	volumeBar.Percent = 100
	volumeBar.Label = fmt.Sprintf("Volume %d%%", volumeBar.Percent)
	volumeBar.BarColor = tui.ColorWhite
	volumeBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	t.volumeBar = volumeBar
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
func (t *TUI) RenderSongInfo() {
	p := widgets.NewParagraph()
	p.Text = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")
	t.p = p
}
func (t *TUI) SetupGrid() {
	grid := tui.NewGrid()
	termWidth, termHeight := tui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)

	grid.Set(
		tui.NewRow(1.6/2,
			tui.NewCol(1.5/2, t.videolist),
			tui.NewCol(0.5/2, t.queuelist)),
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
				//t.HandleSelectedFile(t.filelist.Rows[t.filelist.SelectedRow])
			case "h":
				//t.repo.ShowHiddenFiles = !t.repo.ShowHiddenFiles
				//t.filelist.Rows = t.repo.ListFiles()
				t.RenderUI()
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
			case "<Space>":
				/*
					if t.player != nil {
						go t.player.PauseOrResume()
						if !t.player.Ctrl.Paused {
							t.progressBar.Title = "|| Paused"
						} else {
							t.progressBar.Title = "|> Playing"
						}
						t.RenderUI()
					}
				*/
			case ",":
				/*
					if t.player != nil {
						if t.volumeBar.Percent > 0 {
							t.volumeBar.Percent--
							go t.player.VolumeDown()
							wg.Add(1)
							wg.Done()
						} else {
							t.volumeBar.Percent = 0
						}
						t.volumeBar.Label = fmt.Sprintf("Volume %d%%", t.volumeBar.Percent)
						t.RenderUI()
					}
				*/
			case ".":
				/*
					if t.player != nil {
						if t.volumeBar.Percent < 100 {
							t.volumeBar.Percent++
							go t.player.VolumeUp()
							wg.Add(1)
							wg.Done()
						} else {
							t.volumeBar.Percent = 100
						}
						t.volumeBar.Label = fmt.Sprintf("Volume %d%%", t.volumeBar.Percent)
						t.RenderUI()
					}
				*/
			case "m":
				/*
					if t.player != nil {
						go t.player.Mute()
						wg.Add(1)
						wg.Done()
						if !t.player.Volume.Silent {
							t.volumeBar.Label = "MUTED"
						} else {
							t.volumeBar.Label = fmt.Sprintf("Volume %d%%", t.volumeBar.Percent)
						}
						t.RenderUI()
					}
				*/
			case "<Resize>":
				payload := e.Payload.(tui.Resize)
				t.grid.SetRect(0, 0, payload.Width, payload.Height)
				tui.Clear()
				t.RenderUI()
			}

		case <-*t.ticker:
			t.RenderUI()
		case <-*t.tickerProgresBar:
			/*
				if t.player != nil && !t.player.Ctrl.Paused && t.progressBar.Percent < 100 {
					t.progressBar.Percent++
				}
			*/
			t.RenderUI()
		}
		/*
			if t.player != nil && !t.player.Ctrl.Paused {
				var now = t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second)
				var duration = t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second)
				t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", now, duration)
				t.RenderUI()
			}
		*/
	}
}
func (t *TUI) HandleSelectedFile(filename string) {
	var persistInfo = false

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
	t.progressBar.Label = filename
	//var now = t.player.Samplerate.D(t.player.Streamer.Position()).Round(time.Second)
	//var duration = t.player.Samplerate.D(t.player.Streamer.Len()).Round(time.Second)
	//t.p.Text = fmt.Sprintf("Time:%s  Duration:%s", now, duration)

	if persistInfo {
		//t.player.Volume.Volume = player.OLD_VOLUME
		//t.volumeBar.Percent = player.OLD_PERCENT
	}
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
	go t.RenderQueueList()
	go t.RenderVolumeMixer()
	go t.RenderProgressBar()
	go t.RenderSongInfo()
	wg.Add(5)
	wg.Done()
	time.Sleep(time.Millisecond * 500)
	t.SetupGrid()
	t.uiEvents = tui.PollEvents()
	t.ticker = &time.NewTicker(time.Microsecond).C
	t.tickerProgresBar = &time.NewTicker(time.Second).C
	t.HandleTUIEvents()

}
