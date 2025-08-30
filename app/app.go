package app

import (
	"log"
	"net/http"

	"github.com/eshentials/banking-app/domain"
	"github.com/eshentials/banking-app/service"
	"github.com/gorilla/mux"
)

func StartServer() {

	router := mux.NewRouter()
	repo, err := domain.NewCustomerRepositoryDB()
	if err != nil {
		log.Fatalf("Could not create customer repository: %v", err)
	}
	ch := CustomerHandlers{service: service.NewCustomerService(repo)}
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer)

	log.Fatal(http.ListenAndServe(":8000", router))
}
