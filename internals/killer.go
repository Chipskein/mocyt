package internals

import (
	"log"
	"os"

	"github.com/Chipskein/mocyt/internals/cache_handler"
	"github.com/Chipskein/mocyt/internals/mpv"
)

func KillThemAll() {
	if mpv.CheckIfMpvIsRunning() {
		err := mpv.Stop()
		if err != nil {
			log.Println(err)
		}
		err = os.Remove(mpv.DEFAULT_MPV_SOCKET_PATH)
		if err != nil {
			log.Println(err)
		}
	}

	cache := cache_handler.CheckIfCacheFileExists()
	if cache {
		err := os.Remove(cache_handler.CACHE_FILE_PATH)
		if err != nil {
			log.Println(err)
		}
	}

}
