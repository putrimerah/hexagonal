package app

import (
	"encoding/json"
	"net/http"

	"github.com/irfan/hexagonal/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
