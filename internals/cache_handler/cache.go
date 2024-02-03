package cache_handler

import (
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

const Cache_file_name = "/tmp/cache-ytacli"

func WriteInfo(pf *PlayerInformation) {
	out, err := proto.Marshal(pf)
	if err != nil {
		log.Fatalln("Failed to encode PlayerInformation:", err)
	}
	file, err := os.Create(Cache_file_name)
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
	in, err := os.ReadFile(Cache_file_name)
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
	if _, err := os.Stat(Cache_file_name); os.IsNotExist(err) {
		return false
	}
	return true
}
