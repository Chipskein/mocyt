package utils

import "testing"

func TestParseListString(t *testing.T) {
	id, duration, videoName := ParseListString("[iPbeKLAu-eI] [4m 24s] Persona 5 OST 88 - The Whims of Fate")

	if id != "iPbeKLAu-eI" {
		t.Fatalf("ID %s is not equal to %s\n", id, "iPbeKLAu-eI")
	}
	if videoName != "Persona 5 OST 88 - The Whims of Fate" {
		t.Fatalf("videoName %s is not equal to %s\n", videoName, "Persona 5 OST 88 - The Whims of Fate")
	}
	if duration != "4m 24s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "4m 24s")
	}
}

func TestConvertPTISO8061(t *testing.T) {
	duration := ConvertPTISO8061("PT4M13S")
	if duration != "4m 13s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "4m 13s")
	}
}
func TestConvertSecondsToString(t *testing.T) {
	duration := ConvertSecondsToString(253)
	if duration != "4m 13s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "4m 13s")
	}
	duration = ConvertSecondsToString(7240)
	if duration != "2h 40s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "2h 40s")
	}
}
