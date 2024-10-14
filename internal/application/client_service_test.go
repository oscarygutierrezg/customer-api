package application

import (
	"customer-api/internal/domain"
	"customer-api/internal/ports" // Asegúrate de que la ruta sea correcta
	"errors"
	"testing"
)

// TestGetClient_ValidID_ReturnsClient prueba el caso en que se solicita un customero válido.
func TestGetClient_ValidID_ReturnsClient(t *testing.T) {
	// Given
	mockRepo := &ports.MockClientRepository{}

	customerID := "1"
	expectedClient := &domain.Client{
		ID:   customerID,
		Name: "Laptop",
	}

	mockRepo.On("GetClientByID", customerID).Return(expectedClient, nil)

	customerService := NewClientService(mockRepo)

	// When
	customer, err := customerService.GetClient(customerID)

	// Then
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if customer == nil || customer.ID != expectedClient.ID {
		t.Fatalf("expected customer %+v, got %+v", expectedClient, customer)
	}

	mockRepo.AssertExpectations(t)
}

// TestGetClient_InvalidID_ReturnsError prueba el caso en que se solicita un customero no válido.
func TestGetClient_InvalidID_ReturnsError(t *testing.T) {
	// Given
	mockRepo := &ports.MockClientRepository{}

	invalidID := "1"

	mockRepo.On("GetClientByID", invalidID).Return(nil, errors.New("customer not found"))

	customerService := NewClientService(mockRepo)

	// When
	customer, err := customerService.GetClient(invalidID)

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
