package handlers

import (
	"backend-boilerplate/models"
	"backend-boilerplate/utils"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&input)
	result := utils.DB.Where("username = ?", input.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound || !utils.CheckPasswordHash(input.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
