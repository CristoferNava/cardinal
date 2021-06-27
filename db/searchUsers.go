package db

import (
	"context"
	"log"
	"time"

	"github.com/CristoferNava/cardinal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* SearchUsers get all the users given a restriction, */
func SearchUsers(userID, search, kind string, page int64) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("cardinal")
	usersCollection := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := usersCollection.Find(ctx, query, findOptions)
	if err != nil {
		log.Println(err.Error())
		return results, false
	}

	var found, include bool
	for cursor.Next(ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err.Error())
			return results, false
		}
		var relation models.UserFollowUser
		relation.User1ID = userID
		relation.User2ID = user.ID.Hex()

		include = false
		found, _ = CheckRelation(relation)
		if kind == "new" && !found {
			include = true
		}
		if kind == "follow" && found {
			include = true
		}
		if relation.User2ID == userID {
			include = false
		}
		if include {
			user.Password = ""
			user.Bio = ""
			user.Website = ""
			user.Location = ""
			user.Email = ""
			user.Banner = ""

			results = append(results, &user)
		}
	}
	err = cursor.Err()
	if err != nil {
		log.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
