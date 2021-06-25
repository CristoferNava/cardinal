package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	mw "github.com/CristoferNava/cardinal/middlewares"
	"github.com/CristoferNava/cardinal/routers"
)

// Handle sets the PORT, the Handler with cors and Listen and Serve
func Handle() {
	router := mux.NewRouter()

	router.HandleFunc("/sign-up", mw.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/log-in", mw.CheckDB(routers.LogIn)).Methods("POST")
	router.HandleFunc("/show-profile", mw.CheckDB(mw.ValidateJWT(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/change-profile", mw.CheckDB(mw.ValidateJWT(routers.ChangeProfile))).Methods("PUT")

	router.HandleFunc("/create-tweet", mw.CheckDB(mw.ValidateJWT(routers.CreateTweet))).Methods("POST")
	router.HandleFunc("/show-user-tweets", mw.CheckDB(mw.ValidateJWT(routers.ShowUserTweets))).Methods("GET")
	router.HandleFunc("/remove-tweet", mw.CheckDB(mw.ValidateJWT(routers.RemoveTweet))).Methods("DELETE")

	router.HandleFunc("/upload-avatar", mw.CheckDB(mw.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/serve-avatar", mw.CheckDB(routers.ServeAvatar)).Methods("GET")
	router.HandleFunc("/upload-banner", mw.CheckDB(mw.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/serve-banner", mw.CheckDB(routers.ServeBanner)).Methods("GET")

	router.HandleFunc("/create-relation", mw.CheckDB(mw.ValidateJWT(routers.CreateRelation))).Methods("POST")
	router.HandleFunc("/remove-relation", mw.CheckDB(mw.ValidateJWT(routers.RemoveRelation))).Methods("DELETE")
	router.HandleFunc("/check-relation", mw.CheckDB(mw.ValidateJWT(routers.CheckRelation))).Methods("GET")

	router.HandleFunc("/search-users", mw.CheckDB(mw.ValidateJWT(routers.SearchUsers))).Methods("GET")

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
