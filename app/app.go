package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/irfan/hexagonal/domain"
	"github.com/irfan/hexagonal/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:6677", router))
}
