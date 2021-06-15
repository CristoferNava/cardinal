package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// ChangeProfile receives the request of the client to change a given a profile
func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request format "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = db.ChangeProfile(user, IDUser)
	if err != nil {
		http.Error(w, "Problem with the server "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
