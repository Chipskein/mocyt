package utils

import (
	"strings"
)

func ParseListString(listString string) (id string, videoName string, channelName string, duration string, err error) {
	//EXPECTED INPUT "Persona 5 OST 88 - The Whims of Fate\nID:iPbeKLAu-eI\nDuration:4m 24s\nChannel:Teste\n"
	tmp := strings.Split(listString, "\n")
	videoName = tmp[0]
	id = strings.Split(tmp[1], ":")[1]
	duration = strings.Split(tmp[2], ":")[1]
	channelName = strings.Split(tmp[3], ":")[1]
	return id, videoName, channelName, duration, nil
}
