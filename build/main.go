package main

import (
	"log"

	"github.com/pikastore/pikastore"
)

func main() {
	app := pikastore.New()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
