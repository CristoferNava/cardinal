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

	router.HandleFunc("/sign-up", middlewares.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/show-profile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ShowProfile))).Methods("GEt")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
