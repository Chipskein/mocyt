package youtube

import (
	"chipskein/yta-cli/internals/utils"
	"errors"
	"fmt"

	"google.golang.org/api/youtube/v3"
)

type YoutubeRepository struct {
	Service *youtube.Service
}

type Video struct {
	Id          string
	Title       string
	Duration    string
	ChannelName string
}

func (ytr *YoutubeRepository) ListVideos(searchTxt string) ([]string, error) {
	var videos []string
	service := ytr.Service
	if service == nil {
		return videos, errors.New("nil pointer at services in ytr")
	}
	call := service.Search.List([]string{"snippet"}).Q(searchTxt).Type("video").MaxResults(20)
	response, err := call.Do()
	if err != nil {
		return videos, err
	}
	for _, video := range response.Items {
		var id = video.Id.VideoId
		call := service.Videos.List([]string{"snippet", "contentDetails"}).Id(id)
		response, err := call.Do()
		if err != nil {
			return videos, err
		}
		if len(response.Items) == 0 {
			return videos, fmt.Errorf("could not found any video with id:%s", id)
		}
		var title = response.Items[0].Snippet.Title
		var channelName = response.Items[0].Snippet.ChannelTitle
		var duration = utils.ConvertPTISO8061(response.Items[0].ContentDetails.Duration)
		var video = Video{
			Id:          id,
			Title:       title,
			Duration:    duration,
			ChannelName: channelName}
		var liststring = fmt.Sprintf("[%s] [%s] %s", video.Id, video.Duration, video.Title)
		videos = append(videos, liststring)

	}
	return videos, nil
}
