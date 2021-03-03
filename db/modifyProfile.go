package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyProfile modifies the profile of a user
func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConection.Database("cardinal")
	users := db.Collection("users")

	record := make(map[string]interface{})
	if len(u.Name) > 0 {
		record["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		record["lastName"] = u.LastName
	}
	record["birthday"] = u.Birthdate
	if len(u.Avatar) > 0 {
		record["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		record["banner"] = u.Banner
	}
	if len(u.Bio) > 0 {
		record["bio"] = u.Bio
	}
	if len(u.Location) > 0 {
		record["location"] = u.Location
	}
	if len(u.Web) > 0 {
		record["web"] = u.Web
	}

	updateString := bson.M{"$set": record}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := users.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}
	return true, nil
}
