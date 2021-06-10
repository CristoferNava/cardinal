package db

import (
	"context"
	"log"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile looks up for a given ID user profile in the database
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	users := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}
	err := users.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		log.Println("Profile not found " + err.Error())
		return profile, err
	}
	return profile, nil
}
