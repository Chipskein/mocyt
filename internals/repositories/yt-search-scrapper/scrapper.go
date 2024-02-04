package ytsearchscrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type YoutubeSearch struct {
	SearchTerms string
	MaxResults  int
	Videos      []Video
}

type Video struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}

func NewYoutubeSearch(searchTerms string, maxResults int) *YoutubeSearch {
	return &YoutubeSearch{
		SearchTerms: searchTerms,
		MaxResults:  maxResults,
		Videos:      search(searchTerms, maxResults),
	}
}

func search(searchTerms string, maxResults int) []Video {
	encodedSearch := url.QueryEscape(searchTerms)
	baseURL := "https://youtube.com"
	url := fmt.Sprintf("%s/results?search_query=%s", baseURL, encodedSearch)

	var responseText string
	for !strings.Contains(responseText, "ytInitialData") {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error fetching YouTube search results:", err)
			return nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil
		}

		responseText = string(body)
	}

	results := parseHTML(responseText, maxResults)
	return results
}
func indexAt(s, sep string, n int) int {
	idx := strings.Index(s[n:], sep)
	if idx > -1 {
		idx += n
	}
	return idx
}
func parseHTML(response string, maxResults int) []Video {
	var results []Video
	start := strings.Index(response, "ytInitialData") + len("ytInitialData") + 3
	end := indexAt(response, "};", start) + 1
	jsonStr := response[start:end]
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
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
					video := Video{
						ID:       videoId,
						Title:    title,
						Duration: duration,
					}
					results = append(results, video)
				}
			}

		}

	}
	return results
}
