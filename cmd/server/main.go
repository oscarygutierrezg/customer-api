package main

import (
	"customer-api/internal/adapters/api/v1"
	"customer-api/internal/adapters/repository"
	"customer-api/internal/application"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	repo := repository.NewInMemoryClientRepository()
	customerService := application.NewClientService(repo)
	customerHandler := v1.NewClientHandler(customerService)

	r := mux.NewRouter()

	r.HandleFunc("/v1/customers/{id}", customerHandler.GetClient).Methods("GET")

	log.Println("Server started at :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
