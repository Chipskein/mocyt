package components

import (
	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type VideoList struct {
	root *widgets.List
}

func InitVideoList() *VideoList {
	videolist := widgets.NewList()
	videolist.Rows = []string{"Press '/' to open searchBox and search for a video"}
	videolist.Title = "Video List"
	videolist.TitleStyle.Fg = tui.ColorWhite
	videolist.SelectedRowStyle.Fg = tui.ColorBlack
	videolist.SelectedRowStyle.Bg = tui.ColorWhite
	videolist.TextStyle.Fg = tui.ColorWhite
	return &VideoList{root: videolist}
}
