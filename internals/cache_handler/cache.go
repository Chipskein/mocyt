package cache_handler

import (
	"log"
	"os"
	"path"

	"google.golang.org/protobuf/proto"
)

var Cache_FILE_PATH = path.Join(os.TempDir(), "cache-playback-info-mocyt")

func WriteInfo(pf *PlayerInformation) {
	out, err := proto.Marshal(pf)
	if err != nil {
		log.Fatalln("Failed to encode PlayerInformation:", err)
	}
	file, err := os.Create(Cache_FILE_PATH)
	if err != nil {
		log.Fatalln("Failed to write PlayerInformation:", err)
	}
	_, err = file.Write(out)
	if err != nil {
		log.Fatalln("Failed to write PlayerInformation:", err)
	}
	defer file.Close()
}
func ReadInfo() *PlayerInformation {
	in, err := os.ReadFile(Cache_FILE_PATH)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	info := &PlayerInformation{}
	if err := proto.Unmarshal(in, info); err != nil {
		log.Fatalln("Failed to parse PlayerInformation:", err)
	}
	return info
}
func CheckIfCacheFileExists() bool {
	if _, err := os.Stat(Cache_FILE_PATH); os.IsNotExist(err) {
		return false
	}
	return true
}
