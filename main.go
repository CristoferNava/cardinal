package main

import (
	"log"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/handlers"
)

func main() {
	if !db.CheckConection() {
		log.Fatal("No conection to the DB")
		return
	}
	handlers.Handlers()
}
