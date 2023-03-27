package model

import (
	"portal/server/database"
)

type IModelService struct {}

func ModelService() *IModelService {
	return &IModelService{}
}

func (m *IModelService) Create(opts EModel) (EModel, error) {
	model := EModel{
		Name: opts.Name,
		HuggingFaceID: opts.HuggingFaceID,
	}

	result := database.DB.Create(&model)
	if result.Error != nil {
		return EModel{}, result.Error
	}
	return model, nil
}

// Get a model
func (m *IModelService) Get(id string) (EModel, error) {
	var model EModel
	// Find by ID
	result := database.DB.First(&model, "id = ?", id)
	if result.Error != nil {
		return EModel{}, result.Error
	}
	return model, nil
}

// List all models
func (m *IModelService) List() ([]EModel, error) {
	var models []EModel
	result := database.DB.Find(&models)
	if result.Error != nil {
		return []EModel{}, result.Error
	}
	return models, nil
}

// Update a model
func (m *IModelService) Update(id string, model EModel) (EModel, error) {
	result := database.DB.Model(&model).Where("id = ?", id).Updates(model)
	if result.Error != nil {
		return EModel{}, result.Error
	}
	return model, nil
}
