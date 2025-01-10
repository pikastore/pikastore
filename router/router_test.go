package router_test

import (
	"log"
	"testing"

	"github.com/pikastore/pikastore/router"
)

func TestRouter(t *testing.T) {
	r := router.New(&router.Config{
		PORT: ":3000",
		SSL:  false,
	})
	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
