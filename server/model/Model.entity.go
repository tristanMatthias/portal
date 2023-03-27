package model

import (
	"portal/server/database"
)

type EModel struct {
  database.BaseEntity
  Name string `json:"name"`
  HuggingFaceID string `gorm:"null" json:"huggingFaceID"`
  Downloaded bool `gorm:"default:false" json:"downloaded"`
  DownloadProgress int `gorm:"default:0" json:"downloadProgress"`
  APIKey string `gorm:"null" json:"apiKey"`
}
