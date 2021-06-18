package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTweets inserst a Tweet into the database
func CreateTweet(t models.Tweet) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	tweets := db.Collection("tweets")

	record := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := tweets.InsertOne(ctx, record)
	if err != nil {
		return "", err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), nil
}
