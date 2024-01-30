package components

import (
	"fmt"

	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type VolumeMixer struct {
	Root            *widgets.Gauge
	previus_percent int
}

func (v *VolumeMixer) UpdatePercent(percent int) {
	v.Root.Percent = percent
	v.Root.Label = fmt.Sprintf("Volume %d%%", v.Root.Percent)
}
func (v *VolumeMixer) SetMute(mute bool) {
	if mute {
		v.previus_percent = v.Root.Percent
		v.Root.Percent = 100
		v.Root.Label = "MUTED"
	} else {
		v.Root.Percent = v.previus_percent
		v.Root.Label = fmt.Sprintf("Volume %d%%", v.Root.Percent)
	}
}

func InitVolumeMixer() *VolumeMixer {
	volumeBar := widgets.NewGauge()
	volumeBar.TitleStyle.Fg = tui.ColorWhite
	volumeBar.Percent = 100
	volumeBar.Label = fmt.Sprintf("Volume %d%%", volumeBar.Percent)
	volumeBar.BarColor = tui.ColorWhite
	volumeBar.LabelStyle = tui.NewStyle(tui.ColorWhite)
	return &VolumeMixer{Root: volumeBar, previus_percent: 100}
}
