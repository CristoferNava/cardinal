package routers

import (
	"net/http"

	"github.com/CristoferNava/cardinal/db"
)

// RemoveTweet is the handlers that given an tweetID removes it
func RemoveTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must provide an id", http.StatusBadRequest)
		return
	}
	err := db.RemoveTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "An error ocurred while trying to remove the tweet "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
