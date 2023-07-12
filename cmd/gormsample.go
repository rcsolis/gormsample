package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rcsolis/gormsample/internal/routes"
)

func main() {
	// Create router
	router := mux.NewRouter()
	// Route handlers
	router.HandleFunc("/api/v1/", routes.IndexHandler).Methods("GET")

	// Start server
	http.ListenAndServe(":5000", router)
}
