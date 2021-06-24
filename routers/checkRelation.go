package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// CheckRelation handles the client request for checking if a relation between two users exists
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	userFollowedID := r.URL.Query().Get("id")

	var t models.UserFollowUser
	t.User1ID = IDUser
	t.User2ID = userFollowedID

	var response models.ResponseCheckRelation
	_, err := db.CheckRelation(t)
	if err != nil {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
