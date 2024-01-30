package mpv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

// Conferir JSON IPC do mpv
// https://github.com/mpv-player/mpv/blob/master/DOCS/man/ipc.rst
const DEFAULT_MPV_SOCKET_PATH = "/tmp/mpv-socket"

type IpcJSONMVPResponse struct {
	Data  float32 `json:"data"`
	Error string  `json:"error"`
}

func sendIPCCommand(cmd string) (*IpcJSONMVPResponse, error) {
	var res = &IpcJSONMVPResponse{}
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return res, err
	}
	defer c.Close()
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return res, err
	}
	err = json.NewDecoder(c).Decode(res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func SetSpeed(speed float32) error {
	if speed < 0 || speed > 5 {
		return errors.New("speed value is invalid Should be a Float32 between 0 and 5")
	}
	var cmd = fmt.Sprintf(`{ "command": ["set_property", "speed", %f]}`+"\n", speed)
	_, err := sendIPCCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}
func GetPlayBackTimeMicroSecond() (int, error) {
	var cmd = `{ "command": ["get_property", "playback-time"]}` + "\n"
	res, err := sendIPCCommand(cmd)
	if err != nil {
		return 0, err
	}
	return int(res.Data), nil
}
func GetVolume() (float32, error) {
	var cmd = `{ "command": ["get_property", "volume"]}` + "\n"
	res, err := sendIPCCommand(cmd)
	if err != nil {
		return 0, err
	}
	return res.Data, err
}
func SetVolume(volume int) error {
	if volume < 0 || volume > 100 {
		return errors.New("volume is invalid Should be a int between 0 and 100")
	}
	var cmd = fmt.Sprintf(`{ "command": ["set_property", "volume", %d]}`+"\n", volume)
	_, err := sendIPCCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}
func Stop() error {
	var cmd = `{ "command": ["quit"]}` + "\n"
	_, err := sendIPCCommand(cmd)
	if err != nil {
		return err
	}
	return nil
}
func Pause(pause bool) error {
	var cmd = fmt.Sprintf(`{ "command": ["set_property", "pause", %t]}`+"\n", pause)
	_, err := sendIPCCommand(cmd)
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
