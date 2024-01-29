package mpv

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

// Conferir JSON IPC do mpv
// https://github.com/mpv-player/mpv/blob/master/DOCS/man/ipc.rst
const DEFAULT_MPV_SOCKET_PATH = "/tmp/mpv-socket"

func SetVolume(volume int) error {
	return nil
}
func Stop() error {
	return nil
}
func Pause(pause bool) error {
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return err
	}
	defer c.Close()

	var cmd = fmt.Sprintf(`{ "command": ["set_property", "pause", %t]}`+"\n", pause)
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}
func Play(startDownloadStreamCMD *exec.Cmd, stdin io.ReadCloser) error {
	devnull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	path, err := exec.LookPath("mpv")
	if err != nil {
		return err
	}
	cmdArguments := []string{"-", fmt.Sprintf("-input-ipc-server=%s", DEFAULT_MPV_SOCKET_PATH)}
	playFromStreamCMD := exec.Command(path, cmdArguments...)
	playFromStreamCMD.Stdin = stdin
	playFromStreamCMD.Stdout = devnull
	playFromStreamCMD.Stderr = os.Stderr

	fmt.Println(startDownloadStreamCMD.String())
	fmt.Println(playFromStreamCMD.String())

	err = startDownloadStreamCMD.Start()
	if err != nil {
		return err
	}
	err = playFromStreamCMD.Start()
	if err != nil {
		return err
	}
	err = startDownloadStreamCMD.Wait()
	if err != nil {
		return err
	}
	err = playFromStreamCMD.Wait()
	if err != nil {
		return err
	}
	return nil
}
