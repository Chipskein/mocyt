package protobuf

import (
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

const cache_file_name = "/tmp/cache-ytacli"

func WriteInfo(liststring string, playbackTime int, duration int, percentProgressBar int, volume int, paused bool, playing bool, PID string) {
	teste := &PlayerInformation{
		ListString:        liststring,
		PlaybackTime:      int32(playbackTime),
		Duration:          int32(duration),
		PercentProgresBar: int32(percentProgressBar),
		Volume:            int32(volume),
		Paused:            paused,
		Playing:           playing,
		PidMPV:            PID}
	out, err := proto.Marshal(teste)
	if err != nil {
		log.Fatalln("Failed to encode PlayerInformation:", err)
	}
	file, err := os.OpenFile(cache_file_name, os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		log.Fatalln("Failed to write PlayerInformation:", err)
	}
	_, err = file.Write(out)
	if err != nil {
		log.Fatalln("Failed to write PlayerInformation:", err)
	}
}
func ReadInfo() *PlayerInformation {
	// Read the existing address book.
	in, err := os.ReadFile("testando")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	info := &PlayerInformation{}
	if err := proto.Unmarshal(in, info); err != nil {
		log.Fatalln("Failed to parse PlayerInformation:", err)
	}
	return info
}
