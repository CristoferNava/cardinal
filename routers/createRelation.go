package routers

import (
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// CreateRelation handles the client request to create a UserFollowUser relationship
func CreateRelation(w http.ResponseWriter, r *http.Request) {
	userToFollowID := r.URL.Query().Get("id")
	if len(userToFollowID) < 1 {
		http.Error(w, "Must provide the ID of the user to follow", http.StatusBadRequest)
		return
	}

	var t models.UserFollowUser
	t.User1ID = IDUser
	t.User2ID = userToFollowID

	status, err := db.CreateRelation(t)
	if err != nil || !status {
		http.Error(w, "An error occurred while trying to create the relation in the DB "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
