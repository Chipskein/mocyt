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

// Set PlayBack Speed  Min 0 Max 5
func SetSpeed(speed float32) error {
	if speed < 0 || speed > 5 {
		return errors.New("Speed value is invalid Should be a Float32 between 0 and 5\n")
	}
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return err
	}
	defer c.Close()

	var cmd = fmt.Sprintf(`{ "command": ["set_property", "speed", %f]}`+"\n", speed)
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return err
	}
	//FIX Broken Pipe Error
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
	return nil
}
func GetPlayBackTimeMicroSecond() (int, error) {
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return 0, err
	}
	defer c.Close()
	var cmd = `{ "command": ["get_property", "playback-time"]}` + "\n"
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return 0, err
	}
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
	if err != nil {
		return 0, err
	}
	return int(res.Data), nil
}
func GetVolume() (float32, error) {
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return 0, err
	}
	defer c.Close()

	var cmd = `{ "command": ["get_property", "volume"]}` + "\n"
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return 0, err
	}
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
	if err != nil {
		return 0, err
	}
	return res.Data, err
}
func SetVolume(volume int) error {
	if volume < 0 || volume > 100 {
		return errors.New("Volume is invalid Should be a int between 0 and 100\n")
	}
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return err
	}
	defer c.Close()
	var cmd = fmt.Sprintf(`{ "command": ["set_property", "volume", %d]}`+"\n", volume)
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return err
	}
	//FIX Broken Pipe Error
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
	return nil
}
func Stop() error {
	c, err := net.Dial("unix", DEFAULT_MPV_SOCKET_PATH)
	if err != nil {
		return err
	}
	defer c.Close()

	var cmd = `{ "command": ["quit"]}` + "\n"
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return err
	}
	//FIX Broken Pipe Error
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
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
	//FIX Broken Pipe Error
	var res = &IpcJSONMVPResponse{}
	err = json.NewDecoder(c).Decode(res)
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
