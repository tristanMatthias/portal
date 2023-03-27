package model

import (
	"portal/server/database"
)

type EModel struct {
  database.BaseEntity
  Name string `json:"name" gorm:"unique"`
  HuggingFaceID string `gorm:"null;unique" json:"huggingFaceID"`
  Downloaded bool `gorm:"default:0" json:"downloaded"`
  DownloadProgress int `gorm:"default:0" json:"downloadProgress"`
  APIKey string `gorm:"null" json:"apiKey"`
}
