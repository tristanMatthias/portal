package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseEntity struct {
	gorm.Model
	ID        string         `json:"id" gorm:"primaryKey`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}


func (base *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}
