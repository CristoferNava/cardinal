package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CristoferNava/cardinal/db"
)

// GetTweetsFollowing handles the cliente request to show the tweets of the users we are following
func GetTweetsFollowing(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Invalid page, must send a valid positive integer "+err.Error(), http.StatusBadRequest)
		return
	}

	result, status := db.GetTweetsFollowing(IDUser, page)
	if !status {
		http.Error(w, "Error while reading the Tweets", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
