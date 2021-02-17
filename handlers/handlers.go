package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/CristoferNava/cardinal/middlewares"
	"github.com/CristoferNava/cardinal/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers set the PORT, the handler and listen and serve.
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("sign-up", middlewares.CheckDB(routers.SignUp)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
