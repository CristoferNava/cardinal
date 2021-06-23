package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/CristoferNava/cardinal/db"
)

// ServeAvatar handles the request from a client for a user and response with his avatar
func ServeAvatar(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if len(userID) < 1 {
		http.Error(w, "Must send the user id", http.StatusBadRequest)
		return
	}

	user, err := db.SearchProfile(userID)
	if err != nil {
		http.Error(w, "User not found "+err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.Open("uploads/avatars/" + user.Avatar)
	if err != nil {
		http.Error(w, "Image not found "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error while copying the image "+err.Error(), http.StatusInternalServerError)
	}
}
