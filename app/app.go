package main

import (
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
