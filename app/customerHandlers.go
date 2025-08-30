package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/eshentials/banking-app/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		xml.NewEncoder(w).Encode(customers)
	}

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ")
}

// func CreateCustomers(w http.ResponseWriter, r *http.Request, customer) {
// 	err := json.NewDecoder(r.Body).Decode(&customer)
// 	if err != nil {
// 		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Fprintf(w, "Customer %s from %s with zipcode %d created successfully", customer.Name, customer.City, customer.Zipcode)
// }

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
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
