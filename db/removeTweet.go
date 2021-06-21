package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RemoveTweet removes a tweet given an tweetID and a userID
func RemoveTweet(tweetID, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	tweetsCollection := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(tweetID)
	condition := bson.M{
		"_id":    objID,
		"userID": userID,
	}

	_, err := tweetsCollection.DeleteOne(ctx, condition)
	return err
}
