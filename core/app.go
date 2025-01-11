package core

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Defines the version whatnot
type App struct {
	version string
}

var app = App{
	version: "0",
}

type Config struct {
	Data  string
	DbCon string
}

// Defaults
const (
	DefaultData  string = "/ps_data"
	DefaultDbcon string = "/ps_db"
)

// Creates the App
func NewApp(config Config) *Config {
	if config.DbCon == "" {
		config.DbCon = DefaultDbcon
	}
	if config.Data == "" {
		config.Data = DefaultData
	}
	return &Config{
		Data:  config.Data,
		DbCon: config.DbCon,
	}
}

// Logger
func Format(message, color string) {
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		log.Fatalln("Unable to get caller")
	}
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln("Unable to get cwd")
	}

	fmt.Printf("%s%s\033[0m |  \033[38;5;177m%s:%d \033[0m| \033[0m%s\n", color, time.Now().Format("15:04:05"), file[len(cwd)+1:], line, message)
}

func Log(input string) {
	Format(input, "\033[38;2;30;215;96m")
}
func Warn(input string) {
	Format(input, "\033[38;2;255;215;95m")
}
func Error(input string, panic bool) {
	Format(input, "\033[38;2;219;60;66m")
	if panic {
		os.Exit(0)
	}
}
