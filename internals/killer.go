package internals

import (
	"chipskein/yta-cli/internals/cache_handler"
	"chipskein/yta-cli/internals/mpv"
	"log"
	"os"
)

func KillThemAll() {
	if mpv.CheckIfMpvIsRunning() {
		err := mpv.Stop()
		if err != nil {
			log.Println(err)
		}
		cache := cache_handler.CheckIfCacheFileExists()
		if cache {
			os.Remove(cache_handler.Cache_file_name)
		}
		os.Remove(mpv.DEFAULT_MPV_SOCKET_PATH)
	}
}
