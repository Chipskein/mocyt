package ui

import (
	"chipskein/yta-cli/internals/utils"
	"slices"

	tui "github.com/gizak/termui/v3"
)

func HandleSearchInputEvents(t *TUI, e tui.Event) {
	switch e.Type {
	case tui.KeyboardEvent:
		var char = e.ID
		if char == "<Enter>" {
			var videos = []string{
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST",
				"[iPbeKLAu-eI] Persona 5 OST 88 - The Whims of Fate - Bfr's OST"}
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
		t.UpdateScreen()
	case "q", "<C-c>", "<Escape>":
		return true
	case "<Enter>":
		//Play
		var searchVideoInput = t.grid.Videolist.Root.Rows[t.grid.Videolist.Root.SelectedRow]
		utils.ExtractYoutubeIDFromListString(searchVideoInput)
		//Play Search Text
		t.UpdateScreen()
	case "<Down>", "j":
		t.grid.Videolist.Root.ScrollDown()
		t.UpdateScreen()
	case "<Up>", "k":
		t.grid.Videolist.Root.ScrollUp()
		t.UpdateScreen()
	case "<End>":
		t.grid.Videolist.Root.ScrollBottom()
		t.UpdateScreen()
	case "<Home>":
		t.grid.Videolist.Root.ScrollTop()
		t.UpdateScreen()
	case "<PageDown>":
		t.grid.Videolist.Root.ScrollHalfPageDown()
		t.UpdateScreen()
	case "<PageUp>":
		t.grid.Videolist.Root.ScrollHalfPageUp()
		t.UpdateScreen()
	case "<Space>": //pause
	case ",": //volume up
	case ".": //volume down
	case "m": //mute
	case "<Resize>":
		payload := e.Payload.(tui.Resize)
		t.grid.Root.SetRect(0, 0, payload.Width, payload.Height)
		tui.Clear()
		t.UpdateScreen()
	}
	return false
}
