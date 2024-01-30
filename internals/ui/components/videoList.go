package components

import (
	tui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type VideoList struct {
	Root *widgets.List
}

func (v *VideoList) Update(videos []string, title string) {
	v.Root.Rows = videos
	v.Root.Title = title
}
func (v *VideoList) Reset() {
	v.Root.Rows = []string{"Press '/' to open searchBox and search for a video"}
	v.Root.Title = "Video List"
}
func (v *VideoList) Clean() {
	v.Root.Rows = []string{}
	v.Root.Title = ""
}
func InitVideoList() *VideoList {
	videolist := widgets.NewList()
	videolist.Rows = []string{"Press '/' to open searchBox and search for a video"}
	videolist.Title = "Video List"
	videolist.TitleStyle.Fg = tui.ColorWhite
	videolist.SelectedRowStyle.Fg = tui.ColorBlack
	videolist.SelectedRowStyle.Bg = tui.ColorWhite
	videolist.TextStyle.Fg = tui.ColorWhite
	return &VideoList{Root: videolist}
}
