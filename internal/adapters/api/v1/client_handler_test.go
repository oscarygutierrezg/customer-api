package v1

import (
	"customer-api/internal/adapters/repository"
	"customer-api/internal/application"
	"customer-api/internal/domain"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClient_ValidID_ReturnsClient(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryClientRepository()
	service := application.NewClientService(mockRepository)

	handler := NewClientHandler(service)

	req, err := http.NewRequest("GET", "/v1/customers/b7fc4cb6-6844-4cd0-95fb-f424a3938eb4", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/customers/{id}", handler.GetClient)

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusOK, rr.Code)

	var customer domain.Client
	if err := json.NewDecoder(rr.Body).Decode(&customer); err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	assert.NotNil(t, &customer)
	assert.NotNil(t, &customer.ID)
	assert.NotNil(t, &customer.Name)

}

func TestGetClient_InvalidID_ReturnsError(t *testing.T) {
	// Given
	mockRepository := repository.NewInMemoryClientRepository()
	service := application.NewClientService(mockRepository)

	handler := NewClientHandler(service)

	req, err := http.NewRequest("GET", "/v1/customers/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/v1/customers/{id:[0-9]+}", handler.GetClient)

	// When
	router.ServeHTTP(rr, req)

	// Then
	assert.Equal(t, http.StatusNotFound, rr.Code)

	assert.Equal(t, "customer not found\n", rr.Body.String())
}
