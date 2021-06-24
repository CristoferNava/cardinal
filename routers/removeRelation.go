package routers

import (
	"net/http"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// RemoveRelation handles the client request for giving the ID of the user to unfollow
func RemoveRelation(w http.ResponseWriter, r *http.Request) {
	userFollowed := r.URL.Query().Get("id")
	var t models.UserFollowUser
	t.User1ID = IDUser
	t.User2ID = userFollowed

	status, err := db.RemoveRelation(t)
	if err != nil || !status {
		http.Error(w, "An error ocurred while trying to remove the relation", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
