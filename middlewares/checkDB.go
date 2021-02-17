package middlewares

import (
	"net/http"

	"github.com/CristoferNava/cardinal/db"
)

// CheckDB middleware for the DB status
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConection() {
			http.Error(w, "DB conection lost", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
