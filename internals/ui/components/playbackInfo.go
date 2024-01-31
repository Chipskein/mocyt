package components

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

type PlaybackInfo struct {
	Root *widgets.Paragraph
}

func (p *PlaybackInfo) Reset() {
	if p.Root != nil {
		p.Root.Text = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")
	}
}
func (p *PlaybackInfo) Update(playBackTime string, duration string) {
	if p.Root.Text != "" {
		p.Root.Text = fmt.Sprintf("Time:%s  Duration:%s", playBackTime, duration)
	} else {
		p.Root.Text = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")
	}
}

func InitPlaybackInfo(text string) *PlaybackInfo {
	p := widgets.NewParagraph()
	if text != "" {
		p.Text = text
	} else {
		p.Text = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")
	}
	return &PlaybackInfo{Root: p}
}
