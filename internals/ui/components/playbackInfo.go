package components

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

var defaultText = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")

type PlaybackInfo struct {
	Root *widgets.Paragraph
}

func (p *PlaybackInfo) Reset() {
	if p.Root != nil {
		p.Root.Text = defaultText
	}
}
func (p *PlaybackInfo) Update(text string) {
	if p.Root.Text != "" {
		p.Root.Text = text
	} else {
		p.Root.Text = defaultText
	}
}

func InitPlaybackInfo(text string) *PlaybackInfo {
	p := widgets.NewParagraph()
	if text != "" {
		p.Text = text
	} else {
		p.Text = defaultText
	}
	return &PlaybackInfo{Root: p}
}
