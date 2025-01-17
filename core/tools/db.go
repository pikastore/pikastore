package tools

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)
func initDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open("./ps_db/db.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
	})
}
func GetDB() *gorm.DB {
	initDB()
	return db
}
