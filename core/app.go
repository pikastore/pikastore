package core

import (
	"log"
	"os"

	"github.com/pikastore/pikastore/core/tools"
)

// Defines the version whatnot
type App struct {
	version string
}

var app = App{
	version: "0",
}

type Config struct {
	DataDir string
	DbDir   string
}

// Defaults
const (
	DefaultData  string = "/ps_data"
	DefaultDbcon string = "/ps_db"
)

// Creates the App
func NewApp(config Config) *Config {
	if config.DbDir == "" {
		config.DbDir = DefaultDbcon
	}
	if config.DataDir == "" {
		config.DataDir = DefaultData
	}
	return &Config{
		DataDir: config.DataDir,
		DbDir:   config.DbDir,
	}
}

func Bootstrap(config Config) {
	ensureDir(config.DbDir)
	db := tools.GetDB()
	db.AutoMigrate(
		&tools.User{},
		&tools.Bucket{},
		&tools.Settings{},
	)

	log.Println("Database initialized and migrations applied successfully!")
}
func ensureDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
