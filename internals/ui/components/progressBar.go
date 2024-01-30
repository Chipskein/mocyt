package components

import (
	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ProgressBar struct {
	Root *widgets.Gauge
}

func (p *ProgressBar) Update(percent int, videoName string) {
	p.Root.Percent = percent
	p.Root.Label = videoName
}
func (p *ProgressBar) Reset() {
	p.Root.Percent = 0
	p.Root.Label = " "
}

func InitProgressBar() *ProgressBar {
	progressBar := widgets.NewGauge()
	progressBar.TitleStyle.Fg = tui.ColorWhite
	progressBar.Percent = 0
	progressBar.Label = " "
	progressBar.BarColor = tui.ColorWhite
	progressBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	return &ProgressBar{Root: progressBar}
}
