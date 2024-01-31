package utils

import "testing"

func TestParseListString(t *testing.T) {
	id, videoName, channelName, duration, err := ParseListString("Persona 5 OST 88 - The Whims of Fate\nID:iPbeKLAu-eI\nDuration:4m 24s\nChannel:Teste\n")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if id != "iPbeKLAu-eI" {
		t.Fatalf("ID %s is not equal to %s\n", id, "iPbeKLAu-eI")
	}
	if videoName != "Persona 5 OST 88 - The Whims of Fate" {
		t.Fatalf("ID %s is not equal to %s\n", videoName, "Persona 5 OST 88 - The Whims of Fate")
	}
	if channelName != "Teste" {
		t.Fatalf("ID %s is not equal to %s\n", channelName, "Teste")
	}
	if duration != "4m 24s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "4m 24s")
	}

}
