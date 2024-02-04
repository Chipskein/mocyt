package repositories

type YoutubeRepository interface {
	ListVideos(searchTxt string) ([]string, error)
}

type Video struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
}
