package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
)

// CreateRelations given models.UserFollowUser type insert it into the relations collections
func CreateRelation(t models.UserFollowUser) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	relationsCollection := db.Collection("relations")

	_, err := relationsCollection.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
