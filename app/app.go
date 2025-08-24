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
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.GetAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", router))
}
