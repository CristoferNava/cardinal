package db

import (
	"context"
	"fmt"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile searchs for a specific profile in the database
func SearchProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoConection.Database("cardinal")
	users := db.Collection("users")

	var profile models.User
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := users.FindOne(ctx, condition).Decode(&profile)

	if err != nil {
		fmt.Println("User not found " + err.Error())
	}

	return profile, nil
}
