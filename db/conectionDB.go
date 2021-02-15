package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url = "mongodb+srv://root:elpollo@cardinal.xgsls.mongodb.net/cardinal?retryWrites=true&w=majority"

// MongoConection gets the connection with the DataBase
var MongoConection = ConectDB()
var clientOptions = options.Client().ApplyURI(url)

// ConectDB builds the conection with the database and returns the client
func ConectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Conexión a la DB establecida")
	return client
}

// CheckConection checks the ping with the data base
func CheckConection() bool {
	err := MongoConection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
