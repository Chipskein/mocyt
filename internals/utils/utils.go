package utils

import (
	"strings"
)

func ExtractYoutubeIDFromListString(listString string) (id string, err error) {
	//Expected Input [YOUTUBEID] VIDEONAME - CHANNEL NAME
	tmp := strings.Split(listString, "]")
	//[YOUTUBEID
	tmp2 := tmp[0]
	id = strings.Replace(tmp2, "[", "", 1)
	return id, nil
}
