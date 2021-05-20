package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// SignUp creates a new user in the database
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error with the received data "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Incorrect email", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Incorrect password", 400)
		return
	}

	_, found, _ := db.CheckUserExists(user.Email)
	if found {
		http.Error(w, "The user already exists", 400)
		return
	}

	_, insertStatus, err := db.CreateUser(user)
	if err != nil {
		http.Error(w, "Couldn't insert the user insert the user in the DB "+err.Error(), 500)
		return
	}

	if !insertStatus {
		http.Error(w, "Problem with Mongo while trying to create the user", 500)
	}

	w.WriteHeader(http.StatusCreated)
}
