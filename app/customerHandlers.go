package app

import (
	"encoding/json"
	"net/http"

	"github.com/PrasadR287/banking/service"
	"github.com/gorilla/mux"
)

// Customer DTO - data transfer object
// type Customer struct {
// 	Name    string `json:"full_name" xml:"name" csv:"name"`
// 	City    string `json:"city" xml:"city" csv:"city"`
// 	Zipcode string `json:"zip_code" xml:"xmlzipcode" csv:"csvzipcode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World")
// }

// CustomerHandlers strcut
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Ashish", "New Delhi", "110075"},
	// 	{"Rob", "New Delhi", "110075"},
	// }

	customers, err := ch.service.GetAllCustomer()

	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else if r.Header.Get("Content-Type") == "text/csv" {
	// 	// w.Header().Add("Content-Type", "text/csv")
	// 	// w1 := csv.NewWriter(os.Stdout)
	// 	// w1.WriteAll(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request recieved")
// }
