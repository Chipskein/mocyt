package ytsearchscrapper

import (
	"testing"
)

func TestListVideos(t *testing.T) {
	repo := &YoutubeScrapper{}
	videos, err := repo.ListVideos("nier automata OST chiptune")
	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(videos) == 0 {
		t.Fatalf("%s", "Videos Empty")
	}
}
