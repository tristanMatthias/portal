package model

import (
	"portal/server/lib"
	"time"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (m *Model) ActionDownload(model string) {
	emitUpdate := func(progress Download) {
		rt.EventsEmit(m.ctx, EventDownloadProgress, progress)
	}

	go DownloadModel(model)

	for {
		DownloadsLock.Lock()

		// Check if the download is complete
		downloadCompleted := false
		for _, download := range Downloads {
			emitUpdate(*download)
			if download.Done {
				downloadCompleted = true
			}
		}
		DownloadsLock.Unlock()

		// Close the connection if the download is complete
		if downloadCompleted {
			break
		}

		time.Sleep(1 * time.Second)
	}

}

// An action that lists directories in the models directory
func (m *Model) ActionModelsList() []string {
	mp := lib.ModelPath(nil)
	dirs, err := lib.ListDirs(mp)

	if err != nil {
		return []string{}
	}

	return dirs
}
