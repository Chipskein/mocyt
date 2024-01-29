package ytdlp

import (
	"io"
	"os"
	"os/exec"
)

func DownloadVideo(URL string) (*exec.Cmd, io.ReadCloser, error) {
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return nil, nil, err
	}
	var commandArgs = []string{
		"-f",
		"bestaudio",
		URL,
		"-o",
		"-"}
	cmd := exec.Command(path, commandArgs...)
	cmd.Stderr = os.Stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	return cmd, pipe, nil
}
