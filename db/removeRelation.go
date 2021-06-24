package db

import (
	"context"
	"time"

	"github.com/CristoferNava/cardinal/models"
)

// RemoveRelation removes a relation given the model
func RemoveRelation(t models.UserFollowUser) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	relationsCollection := db.Collection("relations")

	_, err := relationsCollection.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
