package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/jwt"
	"github.com/CristoferNava/cardinal/models"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application-json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid request, check the credentials "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The email is necessary", 400)
		return
	}

	user, status := db.TryLogin(t.Email, t.Password)
	if !status {
		http.Error(w, "Invalid credentials", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Couldn't generate the JWT "+err.Error(), 500)
	}
	response := models.JWT{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
