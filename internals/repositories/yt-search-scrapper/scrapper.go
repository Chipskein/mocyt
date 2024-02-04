package ytsearchscrapper

import (
	"chipskein/yta-cli/internals/repositories"
	"chipskein/yta-cli/internals/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type YoutubeScrapper struct{}

func (ytc *YoutubeScrapper) ListVideos(searchTxt string) (videos []string, err error) {
	videos, err = search(searchTxt)
	if err != nil {
		return []string{}, err
	}
	return videos, nil
}

func search(searchTerms string) (videos []string, err error) {
	encodedSearch := url.QueryEscape(searchTerms)
	baseURL := "https://youtube.com"
	url := fmt.Sprintf("%s/results?search_query=%s", baseURL, encodedSearch)
	videos = []string{}
	var responseText string
	for !strings.Contains(responseText, "ytInitialData") {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching YouTube search results:", err)
			return videos, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return videos, err
		}
		responseText = string(body)
	}
	videos, err = parseHTML(responseText)
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func parseHTML(response string) (videos []string, err error) {
	var results = []string{}
	start := strings.Index(response, "ytInitialData") + len("ytInitialData") + 3
	end := utils.IndexAt(response, "};", start) + 1
	jsonStr := response[start:end]
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return results, err
	}
	contents := data["contents"].(map[string]interface{})
	twoColumnSearchResultsRenderer := contents["twoColumnSearchResultsRenderer"].(map[string]interface{})
	primaryContents := twoColumnSearchResultsRenderer["primaryContents"].(map[string]interface{})
	sectionListRenderer := primaryContents["sectionListRenderer"].(map[string]interface{})
	contents_arr := sectionListRenderer["contents"].([]interface{})
	for _, content := range contents_arr {
		itemSectionRenderer := content.(map[string]interface{})["itemSectionRenderer"]
		if itemSectionRenderer != nil {
			contents_renderer := itemSectionRenderer.(map[string]interface{})["contents"].([]interface{})
			for _, item := range contents_renderer {
				videoRenderer := item.(map[string]interface{})["videoRenderer"]
				if videoRenderer != nil {
					videoMap := videoRenderer.(map[string]interface{})
					videoId := videoMap["videoId"].(string)
					duration := videoMap["lengthText"].(map[string]interface{})["simpleText"].(string)
					titleMap := videoMap["title"].(map[string]interface{})["runs"].([]interface{})[0].(map[string]interface{})
					title := titleMap["text"].(string)
					video := repositories.Video{
						ID:       videoId,
						Title:    title,
						Duration: utils.ConvertHHMMSSToListString(duration),
					}
					var liststring = fmt.Sprintf("[%s] [%s] %s", video.ID, video.Duration, video.Title)
					results = append(results, liststring)
				}
			}
		}
	}
	return results, nil
}
