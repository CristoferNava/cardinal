package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnection has the client for the DB connection
var MongoConnection = connectDB()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful database conection")
	return client
}

// CheckConnection pings to the database to check if is connected
func CheckConnection() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
