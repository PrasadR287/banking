package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start call
func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// Define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
