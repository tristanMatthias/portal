package database

import (
	"portal/server/lib"

	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dbLocation := lib.ConfigPath("portal.db")
	DB, err = gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
