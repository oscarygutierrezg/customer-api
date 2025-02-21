package v1

import (
	"bytes"
	"customer-api/internal/concurrency"
	"customer-api/internal/domain"
	"customer-api/internal/ports"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCustomer_ValidRequest_CreatesCustomer(t *testing.T) {
	mockService := ports.NewMockCustomerService(t)
	pool := concurrency.NewWorkerPool(1, mockService)
	handler := NewCustomerHandler(mockService, pool)

	// Given
	customer := domain.Customer{ID: "1", Name: "Alice", Active: true}
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	mockService.On("CreateCustomer", customer).Return(nil)
	mockService.On("ValidateCustomer", customer.ID).Return(nil)

	rr := httptest.NewRecorder()

	// When
	handler.CreateCustomer(rr, req)

	// Then
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, status)
	}
}

func TestCreateCustomer_InvalidRequest_ReturnsBadRequest(t *testing.T) {
	mockService := ports.NewMockCustomerService(t)
	pool := concurrency.NewWorkerPool(1, mockService)
	handler := NewCustomerHandler(mockService, pool)

	// Given
	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer([]byte("invalid json")))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	// When
	handler.CreateCustomer(rr, req)

	// Then
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status %d but got %d", http.StatusBadRequest, status)
	}
}

func TestCreateCustomer_ServiceError_ReturnsInternalServerError(t *testing.T) {
	mockService := ports.NewMockCustomerService(t)
	pool := concurrency.NewWorkerPool(1, mockService)
	handler := NewCustomerHandler(mockService, pool)

	// Given
	customer := domain.Customer{ID: "4534534", Name: "Alice", Active: true}
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	mockService.On("ValidateCustomer", customer.ID).Return(nil)
	mockService.On("CreateCustomer", customer).Return(errors.New("service error"))

	rr := httptest.NewRecorder()

	// When
	handler.CreateCustomer(rr, req)

	// Then
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status %d but got %d", http.StatusInternalServerError, status)
	}
}

func TestGetCustomer_ValidID_ReturnsCustomer(t *testing.T) {
	mockService := ports.NewMockCustomerService(t)
	pool := concurrency.NewWorkerPool(1, mockService)
	handler := NewCustomerHandler(mockService, pool)

	// Given
	customerID := "1"
	expectedCustomer := &domain.Customer{ID: customerID, Name: "Alice"}

	mockService.On("GetCustomer", customerID).Return(expectedCustomer, nil)

	req, err := http.NewRequest("GET", "/customers?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// When
	handler.GetCustomer(rr, req)

	// Then
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, status)
	}

	var responseCustomer domain.Customer
	json.NewDecoder(rr.Body).Decode(&responseCustomer)

	if responseCustomer.ID != expectedCustomer.ID || responseCustomer.Name != expectedCustomer.Name {
		t.Errorf("Unexpected customer response: %+v", responseCustomer)
	}
}

func TestGetCustomer_InvalidID_ReturnsNotFound(t *testing.T) {
	mockService := ports.NewMockCustomerService(t)
	pool := concurrency.NewWorkerPool(1, mockService)
	handler := NewCustomerHandler(mockService, pool)

	// Given
	invalidID := "1243344423423"

	mockService.On("GetCustomer", invalidID).Return(nil, errors.New("customer not found"))

	req, err := http.NewRequest("GET", "/customers?id="+invalidID, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// When
	handler.GetCustomer(rr, req)

	// Then
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Expected status %d but got %d", http.StatusNotFound, status)
	}
}
