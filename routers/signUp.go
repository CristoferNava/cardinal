package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// SignUp creates in the DB the sign up of a user
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in the stream of data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The user email is necessary", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "The password must be at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserExists(t.Email)
	if found {
		http.Error(w, "The user already exists in the DB", 400)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "There was an error in the sign up "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "It was no possible to save in the database "+err.Error(), 400)
	}

	w.WriteHeader(http.StatusCreated)
}
