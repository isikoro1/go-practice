package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type SearchResponse struct {
	Query string `json:"query"`
	Limit string `json:"limit"`
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	limit := r.URL.Query().Get("limit")

	response := SearchResponse{
		Query: query,
		Limit: limit,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/search", handleSearch)

	log.Println("server started at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
