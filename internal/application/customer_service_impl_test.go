package application

import (
	"customer-api/internal/domain"
	"customer-api/internal/ports"
	"errors"
	"testing"
)

// TestCreateCustomer_ValidCustomer_CreatesSuccessfully prueba el caso en que se crea un cliente válido.
func TestCreateCustomer_ValidCustomer_CreatesSuccessfully(t *testing.T) {
	// Given
	mockRepo := &ports.MockCustomerRepository{}
	newCustomer := domain.Customer{
		ID:     "new-id",
		Name:   "New Customer",
		Active: true,
	}

	mockRepo.On("Save", newCustomer).Return(nil)

	customerService := NewCustomerServiceImpl(mockRepo)

	// When
	err := customerService.CreateCustomer(newCustomer)

	// Then
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockRepo.AssertExpectations(t)
}

// TestCreateCustomer_InvalidCustomer_ReturnsError prueba el caso en que se intenta crear un cliente no válido.
func TestCreateCustomer_InvalidCustomer_ReturnsError(t *testing.T) {
	// Given
	mockRepo := &ports.MockCustomerRepository{}
	invalidCustomer := domain.Customer{
		ID:     "",
		Name:   "",
		Active: false,
	}

	mockRepo.On("Save", invalidCustomer).Return(errors.New("invalid customer"))

	customerService := NewCustomerServiceImpl(mockRepo)

	// When
	err := customerService.CreateCustomer(invalidCustomer)

	// Then
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if err.Error() != "invalid customer" {
		t.Fatalf("expected error 'invalid customer', got %v", err)
	}

	mockRepo.AssertExpectations(t)
}

// TestGetCustomer_ValidID_ReturnsCustomer prueba el caso en que se solicita un customero válido.
func TestGetCustomer_ValidID_ReturnsCustomer(t *testing.T) {
	// Given
	mockRepo := &ports.MockCustomerRepository{}

	customerID := "1"
	expectedCustomer := &domain.Customer{
		ID:   customerID,
		Name: "Laptop",
	}

	mockRepo.On("FindByID", customerID).Return(expectedCustomer, nil)

	customerService := NewCustomerServiceImpl(mockRepo)

	// When
	customer, err := customerService.GetCustomer(customerID)

	// Then
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if customer == nil || customer.ID != expectedCustomer.ID {
		t.Fatalf("expected customer %+v, got %+v", expectedCustomer, customer)
	}

	mockRepo.AssertExpectations(t)
}

// TestGetCustomer_InvalidID_ReturnsError prueba el caso en que se solicita un customero no válido.
func TestGetCustomer_InvalidID_ReturnsError(t *testing.T) {
	// Given
	mockRepo := &ports.MockCustomerRepository{}

	invalidID := "1"

	mockRepo.On("FindByID", invalidID).Return(nil, errors.New("customer not found"))

	customerService := NewCustomerServiceImpl(mockRepo)

	// When
	customer, err := customerService.GetCustomer(invalidID)

	// Then
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if customer != nil {
		t.Fatalf("expected nil customer, got %+v", customer)
	}
	if err.Error() != "customer not found" {
		t.Fatalf("expected error 'customer not found', got %v", err)
	}

	mockRepo.AssertExpectations(t)
}
