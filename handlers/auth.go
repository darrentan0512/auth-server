package handlers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/mux"
	"myauthserver/models"
	"myauthserver/utils"
)

var users = make(map[string]models.User)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)
	users[user.Username] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}