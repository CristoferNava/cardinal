package middlewares

import (
	"net/http"

	"github.com/CristoferNava/cardinal/db"
)

//  CheckDB is a middlware that checks the connection to the database
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "The connection to the database was lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
