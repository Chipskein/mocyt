package ytdlp

import (
	"io"
	"os"
	"os/exec"
)

func DownloadVideo(URL string) (*exec.Cmd, io.ReadCloser, error) {
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	if err != nil {
		return nil, nil, err
	}
	path, err := exec.LookPath("yt-dlp")
	if err != nil {
		return nil, nil, err
	}
	var commandArgs = []string{
		"--quiet",
		"-f",
		"bestaudio",
		URL,
		"-o",
		"-"}
	cmd := exec.Command(path, commandArgs...)

	cmd.Stderr = devnull
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	return cmd, pipe, nil
}
