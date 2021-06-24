package db

import (
	"context"
	"log"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckRelation checks if a relation between two users exists
func CheckRelation(t models.UserFollowUser) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	relationsCollection := db.Collection("relations")

	condition := bson.M{
		"user1ID": t.User1ID,
		"user2ID": t.User2ID,
	}

	var document models.UserFollowUser
	err := relationsCollection.FindOne(ctx, condition).Decode(&document)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
