package router_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/pikastore/pikastore/router"
)

func TestRouter(t *testing.T) {
	router := router.New()

	router.GET("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"message": "Hello, World!"}`)
	}))
}
