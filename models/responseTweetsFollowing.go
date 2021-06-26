package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ResponseTweetFollowing is the response of the Tweets of the users we are following
type ResponseTweetsFollowing struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	User1ID string             `bson:"user1ID,omitempty" json:"user1ID,omitempty"`
	User2ID string             `bson:"user2ID,omitempty" json:"user2ID,omitempty"`
	Tweets  struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}

// Tweets has have the same collection in the DB
