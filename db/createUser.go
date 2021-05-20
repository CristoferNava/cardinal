package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // we cancel the context to avoid stack all the background contexts

	db := MongoConnection.Database("cardinal")
	users := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := users.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
