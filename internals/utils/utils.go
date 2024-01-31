package utils

import (
	"fmt"
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

func ConvertPTISO8061(duration string) string {
	//EXPECTED INPUT "PT4M13S"
	duration = strings.Replace(duration, "PT", "", 1)
	duration = strings.ToLower(duration)
	duration = strings.Replace(duration, "m", "m ", 1)
	duration = strings.Replace(duration, "h", "h ", 1)
	return duration //EXPECTED OUTPUT "4m 13s"
}
func ConvertSecondsToString(seconds int) string {
	var output string = ""
	hours := seconds / 3600
	seconds = seconds % 3600
	minutes := seconds / 60
	seconds = seconds % 60
	if hours > 0 {
		output += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		output += fmt.Sprintf("%dm ", minutes)
	}
	output += fmt.Sprintf("%ds", seconds)
	return output
}
