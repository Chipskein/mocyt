package components

import (
	tui "github.com/gizak/termui/v3"
)

type Grid struct {
	Root        *tui.Grid
	Videolist   *VideoList
	Plabackinfo *PlaybackInfo
	Volumemixer *VolumeMixer
	Progressbar *ProgressBar
	ProgramInfo *ProgramInfo
}

func Init() *Grid {
	var gd = &Grid{}
	grid := tui.NewGrid()
	gd.Root = grid
	termWidth, termHeight := tui.TerminalDimensions()
	gd.Videolist = InitVideoList()
	gd.Plabackinfo = InitPlaybackInfo("")
	gd.Volumemixer = InitVolumeMixer()
	gd.Progressbar = InitProgressBar()
	gd.ProgramInfo = InitProgramInfo()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		tui.NewRow(1.5/2,
			tui.NewCol(1.5/2, gd.Videolist.Root), tui.NewCol(0.5/2, gd.ProgramInfo.Root)),

		tui.NewRow(0.25/2,
			tui.NewCol(1.5/2, gd.Plabackinfo.Root),
			tui.NewCol(0.5/2, gd.Volumemixer.Root)),
		tui.NewRow(0.25/2, gd.Progressbar.Root),
	)
	return gd
}
