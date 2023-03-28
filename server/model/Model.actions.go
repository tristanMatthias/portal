package model

import (
	"fmt"
	ml "portal/server/model/model_lib"

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


	go ml.DownloadModel(hfModelID, func(progress float64) {
		// Update model's download progress
		service.Update(hfModelID, EModel{ DownloadProgress: int(progress), })

		// Update the model in the frontend
		model.DownloadProgress = int(progress * 100)
		if model.DownloadProgress == 100 {
			fmt.Println("Downloaded", hfModelID)
			// Update the model in the frontend
			model.Downloaded = true
			service.Update(hfModelID, model)
		}
		rt.EventsEmit(m.ctx, EventDownloadProgress, model)
	})

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
