package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ChangeProfile use to change the profile of a user
func ChangeProfile(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	users := db.Collection("users")

	// check for the fields that have been updated
	record := make(map[string]interface{})
	if len(user.Name) > 0 {
		record["Name"] = user.Name
	}
	if len(user.LastName) > 0 {
		record["LastName"] = user.LastName
	}
	if len(user.Banner) > 0 {
		record["Banner"] = user.Banner
	}
	record["Birthdate"] = user.Birthdate
	if len(user.Bio) > 0 {
		record["Bio"] = user.Bio
	}
	if len(user.Location) > 0 {
		record["Location"] = user.Location
	}
	if len(user.Website) > 0 {
		record["Website"] = user.Website
	}

	// update the user in the database
	update := bson.M{
		"$set": record,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := users.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}
