package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanran4/banking/domain"
	"github.com/sanran4/banking/service"
)

func Start() {

	router := mux.NewRouter()

	//wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
