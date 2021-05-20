package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists checks if a user exists in the database using his email
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	users := db.Collection("users")

	condition := bson.M{"email": email}
	var result models.User
	err := users.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
