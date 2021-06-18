package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
)

// CreateTweet is the router that given the request creates a Tweet in the database
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var tweeMessage models.TweetMessage

	err := json.NewDecoder(r.Body).Decode(&tweeMessage)
	if err != nil {
		http.Error(w, "Problem while reading the JSON "+err.Error(), http.StatusBadRequest)
		return
	}

	tweet := models.Tweet{
		UserID:  IDUser,
		Message: tweeMessage.Message,
		Date:    time.Now(),
	}

	_, err = db.CreateTweet(tweet)
	if err != nil {
		http.Error(w, "Couldn't insert the Tweet into the DB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
