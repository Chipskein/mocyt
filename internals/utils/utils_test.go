package utils

import "testing"

func TestExtractYoutubeIDFromListString(t *testing.T) {
	id, err := ExtractYoutubeIDFromListString("[YOUTUBEID] VIDEONAME - CHANNEL NAME")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if id != "YOUTUBEID" {
		t.Fatalf("ID %s is not equal to %s\n", id, "YOUTUBEID")
	}
}
