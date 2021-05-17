package main

import (
	"log"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/handlers"
)

func main() {
	if dbConnection := db.CheckConnection(); !dbConnection {
		log.Fatal("Without connection to the DB")
		return
	}
	handlers.Handle()
}
