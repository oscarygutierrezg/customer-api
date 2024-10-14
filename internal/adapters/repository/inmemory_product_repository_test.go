package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewInMemoryClientRepository_CreatesFiveClients
func TestNewInMemoryClientRepository_CreatesSixClients(t *testing.T) {
	// Given
	repo := NewInMemoryClientRepository()

	// When
	actualCount := len(repo.customers)

	// Then
	assert.Equal(t, 6, actualCount, "Expected 6 customers")
}

// TestGetClientByID_ValidID_ReturnsClient
func TestGetClientByID_ValidID_ReturnsClient(t *testing.T) {
	// Given
	repo := NewInMemoryClientRepository()
	customerID := "b7fc4cb6-6844-4cd0-95fb-f424a3938eb4"

	// When
	customer, err := repo.GetClientByID(customerID)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, customer, "Expected a customer, got nil")
	assert.Equal(t, customerID, customer.ID, "Expected customer ID to match")
}

// TestGetClientByID_InvalidID_ReturnsError
func TestGetClientByID_InvalidID_ReturnsError(t *testing.T) {
	// Given
	repo := NewInMemoryClientRepository()
	customerID := "id" // ID que no existe

	// When
	customer, err := repo.GetClientByID(customerID)

	// Then
	assert.Error(t, err, "Expected an error, got none")
	assert.Nil(t, customer, "Expected nil customer, got a customer")
}
