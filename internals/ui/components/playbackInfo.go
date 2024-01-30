package components

import (
	"fmt"

	"github.com/gizak/termui/v3/widgets"
)

var defaultText = fmt.Sprintf("Time:%s  Duration:%s", "0s", "0s")

type PlaybackInfo struct {
	root *widgets.Paragraph
}

func (p *PlaybackInfo) Reset() {
	if p.root != nil {
		p.root.Text = defaultText
	}
}
func (p *PlaybackInfo) Update(text string) {
	if p.root.Text != "" {
		p.root.Text = text
	} else {
		p.root.Text = defaultText
	}
}

func InitPlaybackInfo(text string) *PlaybackInfo {
	p := widgets.NewParagraph()
	if text != "" {
		p.Text = text
	} else {
		p.Text = defaultText
	}
	return &PlaybackInfo{root: p}
}
