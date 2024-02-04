package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseListString(liststring string) (id string, duration string, videoName string) {
	//EXPECTED INPUT "[iPbeKLAu-eI] [4m 24s] Persona 5 OST 88 - The Whims of Fate"
	var OpenSquareBracketIndex = 0
	var CloseSquareBracketIndex = strings.Index(liststring, "]")
	id = liststring[OpenSquareBracketIndex+1 : CloseSquareBracketIndex]
	liststring = liststring[CloseSquareBracketIndex+2:]

	OpenSquareBracketIndex = 0
	CloseSquareBracketIndex = strings.Index(liststring, "]")
	duration = liststring[OpenSquareBracketIndex+1 : CloseSquareBracketIndex]

	videoName = liststring[CloseSquareBracketIndex+2:]
	return id, duration, videoName
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
func ConvertStringToSeconds(duration string) int {
	//EXPECTED INPUT "4m 13s"
	var seconds int = 0
	tmp_slice := strings.Split(duration, " ")
	for _, t := range tmp_slice {
		if strings.Contains(t, "h") {
			var h int
			fmt.Sscanf(t, "%dh", &h)
			seconds += h * 3600
		} else if strings.Contains(t, "m") {
			var m int
			fmt.Sscanf(t, "%dm", &m)
			seconds += m * 60
		} else if strings.Contains(t, "s") {
			var s int
			fmt.Sscanf(t, "%ds", &s)
			seconds += s
		}
	}
	return seconds
}

func IndexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}

func ConvertHHMMSSToListString(t string) string {
	//input 03:30
	var duration string = ""
	nDoublePoint := strings.Count(t, ":")
	tmp := strings.Split(t, ":")
	switch nDoublePoint {
	case 2:
		hh, _ := strconv.Atoi(tmp[0])
		mm, _ := strconv.Atoi(tmp[1])
		ss, _ := strconv.Atoi(tmp[2])
		duration = fmt.Sprintf("%dh %dm %ds", hh, mm, ss)
	case 1: //mm:ss
		mm, _ := strconv.Atoi(tmp[0])
		ss, _ := strconv.Atoi(tmp[1])
		duration = fmt.Sprintf("%dm %ds", mm, ss)
	}
	//output 3m 30s
	return duration
}
