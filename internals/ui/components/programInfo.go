package components

import (
	"github.com/gizak/termui/v3/widgets"
)

type ProgramInfo struct {
	Root *widgets.Paragraph
}

func InitProgramInfo() *ProgramInfo {
	p := widgets.NewParagraph()
	p.Text = "Basic usage:\n*Play:<Enter>\n*VolumeUp:<.>\n*VolumeDown:<,>\n*Mute:<m>\n*Open Search:</>\n*Search:<Enter>\n*Exit Search:<q>|<Esc>\n*Exit:<q>|<Esc>\n*Clear Search:<Del>\n*Move:<↑,↓>\n"
	return &ProgramInfo{Root: p}
}
