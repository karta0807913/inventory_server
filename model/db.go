package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteDB(dbName string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func InitDB(db *gorm.DB) {
	db.AutoMigrate(&ItemTable{})
}
