package components

import (
	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ProgressBar struct {
	Root *widgets.Gauge
}

func (p *ProgressBar) Update(percent int, videoName string, playing bool) {
	p.Root.Percent = percent
	p.Root.Label = videoName
	if playing {
		p.Root.Title = "|> Playing"
	} else {
		p.Root.Title = "|| Paused"
	}
}
func (p *ProgressBar) Reset() {
	p.Root.Percent = 0
	p.Root.Label = " "
}

func InitProgressBar() *ProgressBar {
	progressBar := widgets.NewGauge()
	progressBar.Title = "|> Welcome to mocyt"
	progressBar.TitleStyle.Fg = tui.ColorWhite
	progressBar.Percent = 0
	progressBar.Label = " "
	progressBar.BarColor = tui.ColorWhite
	progressBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	return &ProgressBar{Root: progressBar}
}
