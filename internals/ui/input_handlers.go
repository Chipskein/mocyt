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
			var videos, _ = t.repository.ListVideos(t.searchTxt)
			t.searchTxt = ""
			t.grid.Videolist.Update(videos, "Press Enter to Play")
			t.UpdateScreen()
			t.shouldRenderSearchBar = false
			break
		}
		if char == "<Space>" {
			char = " "
		}
		if char != "<Backspace>" {
			var invalidChars = []string{"<Up>", "<Down>", "<Left>", "<Right>", "<Insert>", "<Delete>", "<Home>", "<End>", "<Previous>", "<Next>", "<Tab>"}
			isInvalid := slices.Contains(invalidChars, char)
			if isInvalid {
				break
			}
			t.searchTxt += char
		} else {
			if len(t.searchTxt) > 0 {
				t.searchTxt = t.searchTxt[:len(t.searchTxt)-1]
			}
		}

		t.grid.Videolist.Update([]string{t.searchTxt}, "Press Enter to Search")
		t.UpdateScreen()
	}
}

func HandleUserCommands(t *TUI, e tui.Event) (shouldExit bool) {
	switch e.ID {
	case "/":
		t.shouldRenderSearchBar = true
		t.grid.Videolist.Clean()
	case "q", "<C-c>", "<Escape>":
		cache_handler.WriteInfo(t.Current_player_info)
		return true
	case "<Enter>":
		var searchVideoInput = t.grid.Videolist.Root.Rows[t.grid.Videolist.Root.SelectedRow]
		if searchVideoInput != "Press '/' to open searchBox and search for a video" {
			HandleSelectedVideo(t, searchVideoInput)
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
