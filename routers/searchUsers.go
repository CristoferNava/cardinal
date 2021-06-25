package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/CristoferNava/cardinal/db"
)

// SearchUsers handles the cliente request for searching users given a tupe, a page and a regex for search
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	log.Println("todo cool hasta aquí")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Must send a valid number page "+err.Error(), http.StatusBadRequest)
		return
	}

	page64 := int64(pageTemp)
	result, status := db.SearchUsers(IDUser, search, typeUser, page64)
	if !status {
		http.Error(w, "Problem while searching the users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
