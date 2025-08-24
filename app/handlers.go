package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/eshentials/banking-app/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{Name: "John", City: "New York", Zipcode: 10001},
	// 	{Name: "Jane", City: "Los Angeles", Zipcode: 90001},
	// 	{Name: "Mike", City: "Chicago", Zipcode: 60601},
	// }
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ")
}

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zipcode" xml:"zipcode"`
}

func CreateCustomers(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Customer %s from %s with zipcode %d created successfully", customer.Name, customer.City, customer.Zipcode)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//fmt.Fprint(w, vars["customer_id"])
}
