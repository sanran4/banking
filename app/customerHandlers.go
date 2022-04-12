package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sanran4/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponce(w, err.Code, err.AsMessage())
	} else {
		writeResponce(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponce(w, err.Code, err.AsMessage())
	} else {
		writeResponce(w, http.StatusOK, customer)
	}

}

func writeResponce(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
