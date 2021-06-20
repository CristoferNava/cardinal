package db

import (
	"context"
	"log"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ShowUserTweets returns the tweets for a given user and a given page
func ShowUserTweets(userID string, page int64) ([]*models.UserTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("cardinal")
	tweetsCollection := db.Collection("tweets")

	var userTweets []*models.UserTweet
	condition := bson.M{
		"userID": userID,
	}

	ops := options.Find()
	ops.SetLimit(20)                              // we only want to query 20 documents of the collection
	ops.SetSort(bson.D{{Key: "date", Value: -1}}) // inverse order by date (most recent first)
	ops.SetSkip((page - 1) * 20)                  // given a page this control the documents queried (skiped)

	cursor, err := tweetsCollection.Find(ctx, condition, ops)
	if err != nil {
		log.Fatal(err.Error())
		return userTweets, false
	}

	for cursor.Next(context.TODO()) {
		var userTweet models.UserTweet
		err := cursor.Decode(&userTweet)
		if err != nil {
			log.Fatal(err.Error())
			return userTweets, false
		}
		userTweets = append(userTweets, &userTweet)
	}
	return userTweets, true
}
