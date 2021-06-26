package db

import (
	"context"
	"log"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTweetsFollowing(ID string, page int) ([]models.ResponseTweetsFollowing, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	relationsColl := db.Collection("relations")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"user1ID": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "user2ID",
			"foreignField": "userID",
			"as":           "tweets",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweets"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, _ := relationsColl.Aggregate(ctx, conditions)
	var result []models.ResponseTweetsFollowing
	err := cursor.All(ctx, &result)
	if err != nil {
		log.Println(err.Error())
		return result, false
	}
	return result, true
}
