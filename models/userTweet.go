package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserTweet gets a tweet for a given user
type UserTweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID  string             `bson:"userID,omitempty" json:"userID,omitempty"`
	Message string             `bson:"message,omitempty" json:"message,omitempty"`
	Date    time.Time          `bson:"date,omitiempty" json:"date,omitempty"`
}
