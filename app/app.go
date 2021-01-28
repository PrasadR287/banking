package app

import (
	"log"
	"net/http"
)

// Start call
func Start() {

	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
