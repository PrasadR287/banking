package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer DTO - data transfer object
type Customer struct {
	Name    string `json:"full_name" xml:"name" csv:"name"`
	City    string `json:"city" xml:"city" csv:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode" csv:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
