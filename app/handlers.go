package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{name: "John", City: "New York", Zipcode: 10001},
		{name: "Jane", City: "Los Angeles", Zipcode: 90001},
		{name: "Mike", City: "Chicago", Zipcode: 60601},
	}
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
	name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zipcode" xml:"zipcode"`
}
