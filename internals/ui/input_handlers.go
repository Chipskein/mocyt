package ui

import (
	"chipskein/yta-cli/internals/cache_handler"
	"slices"

	tui "github.com/gizak/termui/v3"
)

func HandleSearchInputEvents(t *TUI, e tui.Event) {
	switch e.Type {
	case tui.KeyboardEvent:
		var char = e.ID
		if char == "<Enter>" {
			var videos, err = t.repository.ListVideos(t.Current_player_info.SearchTxt)
			if err != nil {
				panic(err)
			}
			t.Current_player_info.SearchResults = videos
			videos = []string{t.Current_player_info.SearchTxt}
			videos = append(videos, t.Current_player_info.SearchResults...)
			t.grid.Videolist.Update(videos, "Welcome!")
			t.UpdateScreen()
			t.shouldRenderSearchBar = false
			break
		}
		if char == "<Space>" {
			char = " "
		}
		if char == "<Delete>" {
			t.Current_player_info.SearchTxt = ""
			t.grid.Videolist.Root.Rows[0] = t.Current_player_info.SearchTxt
			t.grid.Videolist.Update(t.grid.Videolist.Root.Rows, "Welcome!")
			t.UpdateScreen()
		}
		if char == "<Escape>" || char == "q" {
			t.shouldRenderSearchBar = false
			break
		}

		if char != "<Backspace>" {
			var invalidChars = []string{"<Up>", "<Down>", "<Left>", "<Right>", "<Insert>", "<Delete>", "<Home>", "<End>", "<Previous>", "<Next>", "<Tab>"}
			isInvalid := slices.Contains(invalidChars, char)
			if isInvalid {
				break
			}
			t.Current_player_info.SearchTxt += char
		} else {
			if len(t.Current_player_info.SearchTxt) > 0 {
				t.Current_player_info.SearchTxt = t.Current_player_info.SearchTxt[:len(t.Current_player_info.SearchTxt)-1]
			}
		}
		t.grid.Videolist.Root.Rows[0] = t.Current_player_info.SearchTxt
		t.grid.Videolist.Update(t.grid.Videolist.Root.Rows, "Welcome!")
		t.UpdateScreen()
	}
}

func HandleUserCommands(t *TUI, e tui.Event) (shouldExit bool) {
	switch e.ID {
	case "/":
		t.grid.Videolist.Root.SelectedRow = 0
		t.shouldRenderSearchBar = true
	case "q", "<C-c>", "<Escape>":
		cache_handler.WriteInfo(t.Current_player_info)
		return true
	case "<Enter>":
		var searchVideoInput = t.grid.Videolist.Root.Rows[t.grid.Videolist.Root.SelectedRow]
		if t.grid.Videolist.Root.SelectedRow > 0 {
			HandleSelectedVideo(t, searchVideoInput)
		} else {
			t.grid.Videolist.Root.SelectedRow = 0
			t.shouldRenderSearchBar = true
		}
	case "<Down>", "j":
		t.grid.Videolist.Root.ScrollDown()
	case "<Up>", "k":
		t.grid.Videolist.Root.ScrollUp()
	case "<End>":
		t.grid.Videolist.Root.ScrollBottom()
	case "<Home>":
		t.grid.Videolist.Root.ScrollTop()
	case "<PageDown>":
		t.grid.Videolist.Root.ScrollHalfPageDown()
	case "<PageUp>":
		t.grid.Videolist.Root.ScrollHalfPageUp()
	case "<Space>":
		handlePause(t)
	case ",":
		handleVolumeDown(t)
	case ".":
		handleVolumeUp(t)
	case "m":
		handleMute(t)
	case "<Resize>":
		payload := e.Payload.(tui.Resize)
		t.grid.Root.SetRect(0, 0, payload.Width, payload.Height)
		tui.Clear()
	}
	t.UpdateScreen()
	return false
}
