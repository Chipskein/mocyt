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

func TestConvertStringToSeconds(t *testing.T) {
	duration := ConvertStringToSeconds("4m 13s")
	if duration != 253 {
		t.Fatalf("Duration %d is not equal to %s\n", duration, "4m 13s")
	}
	duration = ConvertStringToSeconds("2h 40s")
	if duration != 7240 {
		t.Fatalf("Duration %d is not equal to %s\n", duration, "2h 40s")
	}
}

func TestConvertHHMMSSToListString(t *testing.T) {
	duration := ConvertHHMMSSToListString("03:30")
	if duration != "3m 30s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "3m 30s")
	}
	duration = ConvertHHMMSSToListString("02:03:30")
	if duration != "2h 3m 30s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "2h 3m 30s")
	}

	duration = ConvertHHMMSSToListString("02:00:30")
	if duration != "2h 0m 30s" {
		t.Fatalf("Duration %s is not equal to %s\n", duration, "2h 0m 30s")
	}
}
