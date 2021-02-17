package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser inserts a new user in the DB
func InsertUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // downs the WithTime in the context

	db := MongoConection.Database("cardinal")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
