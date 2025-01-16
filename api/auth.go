package api

import (
	"net/http"

	"github.com/pikastore/pikastore/router"
)

func RegisterAuth(r *router.Group) {
	r.Post("/auth/register", http.HandlerFunc(RegisterHandler))
	r.Post("/auth/login", http.HandlerFunc(RegisterHandler))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
