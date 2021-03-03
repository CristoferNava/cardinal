package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// UpdateProfile updates the profile of a user
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "The data is incorrect "+err.Error(), 400)
		return
	}

	var status bool
	status, err = db.ModifyProfile(user, UserID)
	if err != nil {
		http.Error(w, "An error trying to update the profile, please try again", 400)
		return
	}

	if status == false {
		http.Error(w, "It was not posible to update the user profile", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
