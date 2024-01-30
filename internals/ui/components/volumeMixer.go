package components

import (
	"fmt"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type VolumeMixer struct {
	root            *widgets.Gauge
	previus_percent int
}

func (v *VolumeMixer) UpdatePercent(percent int) {
	v.root.Percent = percent
	v.root.Label = fmt.Sprintf("Volume %d%%", v.root.Percent)
}
func (v *VolumeMixer) SetMute(mute bool) {
	if mute {
		v.previus_percent = v.root.Percent
		v.root.Percent = 100
		v.root.Label = "MUTED"
	} else {
		v.root.Percent = v.previus_percent
		v.root.Label = fmt.Sprintf("Volume %d%%", v.root.Percent)
	}
}

func InitVolumeMixer() *VolumeMixer {
	volumeBar := widgets.NewGauge()
	volumeBar.TitleStyle.Fg = tui.ColorWhite
	volumeBar.Percent = 100
	volumeBar.Label = fmt.Sprintf("Volume %d%%", volumeBar.Percent)
	volumeBar.BarColor = tui.ColorWhite
	volumeBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	return &VolumeMixer{root: volumeBar, previus_percent: 100}
}
