package api

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pikastore/pikastore/core/tools"
	"github.com/pikastore/pikastore/router"
)

func RegisterAuth(r *router.Group) {
	r.Post("/auth/register", http.HandlerFunc(RegisterHandler))
	r.Post("/auth/login", http.HandlerFunc(RegisterHandler))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//I'm too lazy to try and figure out how to use the schema so this should do for now.
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	db := tools.GetDB()
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if len(input.Username) < 3 || len(input.Password) < 6 {
		http.Error(w, "Username must be at least 3 characters and password at least 6 characters", http.StatusBadRequest)
		return
	}
	var existingUserCount int64
	if err := db.Model(&tools.User{}).Count(&existingUserCount).Error; err != nil {
		http.Error(w, "Failed to check user count", http.StatusInternalServerError)
		return
	}
	if existingUserCount > 0 {
		http.Error(w, "Only one user is allowed", http.StatusConflict)
		return
	}

	h := sha256.Sum256([]byte(input.Password))

	newUser := tools.User{
		Username: input.Username,
		Password: fmt.Sprintf("%x", h),
	}
	if err := db.Create(&newUser).Error; err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
