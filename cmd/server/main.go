package main

import (
	"customer-api/internal/adapters/api/v1"
	"customer-api/internal/adapters/db/repository"
	"customer-api/internal/application"
	"customer-api/internal/concurrency"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewInMemoryCustomerRepository()
	service := application.NewCustomerServiceImpl(repo)
	workerPool := concurrency.NewWorkerPool(5, service)
	customerHandler := v1.NewCustomerHandler(service, workerPool)

	r := mux.NewRouter()

	r.HandleFunc("/v1/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/v1/customers/{id}", customerHandler.GetCustomer).Methods("GET")

	log.Println("Server started at :8083")
	log.Fatal(http.ListenAndServe(":8083", r))
}
