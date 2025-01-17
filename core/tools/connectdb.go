package tools

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		return nil, err
	}
	return db, nil
}
