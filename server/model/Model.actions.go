package model

import (
	ml "portal/server/model/model_lib"
	"time"

	rt "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (m *Model) ActionDownload(hfModelID string) {
	service := ModelService()

	// Check if model is already downloaded
	model, err := service.Get(hfModelID)
	if err == nil && model.Downloaded {
		return
	} else {
		model, _ = service.Create(EModel{
			Name: hfModelID,
			HuggingFaceID: hfModelID,
		})
	}

	emitUpdate := func(progress ml.Download) {
		rt.EventsEmit(m.ctx, EventDownloadProgress, progress)
	}

	go ml.DownloadModel(hfModelID)

	for {
		ml.DownloadsLock.Lock()

		// Check if the download is complete
		downloadCompleted := false
		for _, download := range ml.Downloads {
			emitUpdate(*download)
			if download.Done {
				downloadCompleted = true
				println("Download complete")
				// Update model's download progress
				service.Update(hfModelID, EModel{
					Downloaded:       true,
					DownloadProgress: 100,
				})
			} else {
				println("Download progress:", int(download.Progress * 100 / download.Total))
				// Update model's download progress
				service.Update(hfModelID, EModel{
					DownloadProgress: int(download.Progress * 100 / download.Total),
				})
			}
		}
		ml.DownloadsLock.Unlock()

		// Close the connection if the download is complete
		if downloadCompleted {
			break
		}

		time.Sleep(1 * time.Second)
	}

}

// An action that lists directories in the models directory
func (m *Model) ActionModelsList() ([]EModel, error) {
	service := ModelService()
	models, err := service.List()
	if err != nil {
		return nil, err
	}
	return models, nil
}
