package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteDB(dbName string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func InitDB(db *gorm.DB) error {
	err := db.AutoMigrate(&ItemTable{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&UserData{})
	if err != nil {
		return err
	}
	return nil
}
