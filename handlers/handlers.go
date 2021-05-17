package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handle sets the PORT, the Handler with cors and Listen and Serve
func Handle() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	err := http.ListenAndServe(":"+PORT, handler)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	log.Println("Listening on PORT 8080")
}
