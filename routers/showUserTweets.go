package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CristoferNava/cardinal/db"
)

// ShowUserTweets handle the client request for getting the tweets
func ShowUserTweets(w http.ResponseWriter, r *http.Request) {
	// al the data we get by URL is a string type
	userID := r.URL.Query().Get("id")
	if len(userID) < 1 {
		http.Error(w, "You must send an userID", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send a number of page", http.StatusBadRequest)
		return
	}

	pageInt, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send a number of page (no a letter)", http.StatusBadRequest)
		return
	}

	page := int64(pageInt)
	userTweets, status := db.ShowUserTweets(userID, page)
	if !status {
		http.Error(w, "Problem while trying to get the user tweets", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userTweets)
}
