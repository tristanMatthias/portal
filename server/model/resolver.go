package model

import (
	"context"
	"time"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func OnDownload(ctx context.Context) func(...interface{}) {
	emitUpdate := func(progress Download) {
		rt.EventsEmit(ctx, EventDownloadProgress, progress)
	}

	return func(optionalData ...interface{}) {
		if len(optionalData) > 0 {
			data, ok := optionalData[0].(string)
			if ok {
				go DownloadModel(data)

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
		}
	}
}
