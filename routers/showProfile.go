package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
)

// ShowProfile shows the profile of a given user
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The ID of user is not valid", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)
	if err != nil {
		http.Error(w, "User profile not found"+err.Error(), 400)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
