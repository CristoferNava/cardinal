package models

import "time"

// Tweet struct for saving the data of the T
type Tweet struct {
	UserID  string    `bson:"userID,omitempty" json:"userID,omitempty"`
	Message string    `bson:"message,omitempty" json:"message,omitempty"`
	Date    time.Time `bson:"date,omitempty" json:"date,omitempty"`
}

// TweetMessage is struct that decodes the message from the client request
type TweetMessage struct {
	Message string `bson:"message,omitempty" json:"message,omitempty"`
}
