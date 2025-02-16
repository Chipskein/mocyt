package ytsearchscrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/Chipskein/mocyt/internals/repositories"
	"github.com/Chipskein/mocyt/internals/utils"
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
		return videos, nil
	}
	return videos, nil
}

func parseHTML(response string) (videos []string, err error) {
	var results = []string{}
	startFound := strings.Contains(response, "ytInitialData")
	if !startFound {
		fmt.Println("ytInitialData not found in response")
		return results, fmt.Errorf("ytInitialData not found in response")
	}
	start := strings.Index(response, "ytInitialData") + len("ytInitialData") + 3
	end := utils.IndexAt(response, "};", start) + 1
	jsonStr := response[start:end]
	var data map[string]any
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return results, err
	}
	contents := data["contents"].(map[string]any)
	if contents == nil {
		fmt.Println("contents not found in JSON")
		return results, fmt.Errorf("contents not found in JSON")
	}

	twoColumnSearchResultsRenderer := contents["twoColumnSearchResultsRenderer"].(map[string]any)
	if twoColumnSearchResultsRenderer == nil {
		fmt.Println("twoColumnSearchResultsRenderer not found in JSON")
		return results, fmt.Errorf("twoColumnSearchResultsRenderer not found in JSON")
	}

	primaryContents := twoColumnSearchResultsRenderer["primaryContents"].(map[string]any)
	if primaryContents == nil {
		fmt.Println("primaryContents not found in JSON")
		return results, fmt.Errorf("primaryContents not found in JSON")
	}

	sectionListRenderer := primaryContents["sectionListRenderer"].(map[string]any)
	if sectionListRenderer == nil {
		fmt.Println("sectionListRenderer not found in JSON")
		return results, fmt.Errorf("sectionListRenderer not found in JSON")
	}

	contents_arr := sectionListRenderer["contents"].([]any)
	if contents_arr == nil {
		fmt.Println("contents not found in JSON")
		return results, fmt.Errorf("contents_arr not found in JSON")
	}

	for _, content := range contents_arr {
		itemSectionRenderer := content.(map[string]any)["itemSectionRenderer"]
		if itemSectionRenderer != nil {
			contents_renderer := itemSectionRenderer.(map[string]any)["contents"].([]any)
			if contents_renderer == nil {
				fmt.Println("contents_renderer not found in JSON")
				continue
			}
			for _, item := range contents_renderer {
				if item == nil {
					fmt.Println("item not found in JSON")
					continue
				}
				videoRenderer := item.(map[string]any)["videoRenderer"]
				if videoRenderer != nil {
					videoMap := videoRenderer.(map[string]any)
					if videoMap == nil {
						fmt.Println("videoMap not found in JSON")
						continue
					}

					if videoMap["videoId"] == nil {
						fmt.Println("videoId not found in JSON")
						continue
					}

					videoId := videoMap["videoId"].(string)
					if videoMap["lengthText"] == nil {
						fmt.Println("lengthText not found in JSON")
						continue
					}

					lengthTxtMap := videoMap["lengthText"].(map[string]any)
					if lengthTxtMap == nil {
						fmt.Println("lengthTxtMap not found in JSON")
						continue
					}

					if lengthTxtMap["simpleText"] == nil {
						fmt.Println("simpleText not found in JSON")
						continue
					}
					duration := lengthTxtMap["simpleText"].(string)

					if videoMap["title"] == nil {
						fmt.Println("title not found in JSON")
						continue
					}
					titleMap1 := videoMap["title"].(map[string]any)
					if titleMap1["runs"] == nil {
						fmt.Println("runs not found in JSON")
						continue
					}
					titleRunArr := titleMap1["runs"].([]any)
					if len(titleRunArr) == 0 {
						fmt.Println("titleRunArr not found in JSON")
						continue
					}
					titleMap := titleRunArr[0].(map[string]any)
					if titleMap["text"] == nil {
						fmt.Println("text not found in JSON")
						continue
					}
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
