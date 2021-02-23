package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
	"github.com/CristoferNava/jwt"
)

// Login handles the frontend request for the login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User or password don't valid"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User email is required", 400)
	}

	user, found := db.TryLogin(t.Email, t.Password)
	if found == false {
		http.Error(w, "User or password incorrect")
		return
	}

	// Login was success, so we build the web token
	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "An error occurred when trying to generate the token"+err.Error(), 400)
	}
	response := models.ResponseLogin{Token: jwtKey}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
