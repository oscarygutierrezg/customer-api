package repository

import (
	"customer-api/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewInMemoryCustomerRepository_CreatesSixCustomers
func TestNewInMemoryCustomerRepository_CreatesSixCustomers(t *testing.T) {
	// Given
	repo := NewInMemoryCustomerRepository()

	// When
	actualCount := len(repo.customers)

	// Then
	assert.Equal(t, 6, actualCount, "Expected 6 customers")
}

// TestGetCustomerByID_ValidID_ReturnsCustomer
func TestGetCustomerByID_ValidID_ReturnsCustomer(t *testing.T) {
	// Given
	repo := NewInMemoryCustomerRepository()
	customerID := "b7fc4cb6-6844-4cd0-95fb-f424a3938eb4"

	// When
	customer, err := repo.FindByID(customerID)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, customer, "Expected a customer, got nil")
	assert.Equal(t, customerID, customer.ID, "Expected customer ID to match")
}

// TestGetCustomerByID_InvalidID_ReturnsError
func TestGetCustomerByID_InvalidID_ReturnsError(t *testing.T) {
	// Given
	repo := NewInMemoryCustomerRepository()
	customerID := "id" // ID que no existe

	// When
	customer, err := repo.FindByID(customerID)

	// Then
	assert.Error(t, err, "Expected an error, got none")
	assert.Nil(t, customer, "Expected nil customer, got a customer")
}

// TestSave_ValidCustomer_SavesSuccessfully
func TestSave_ValidCustomer_SavesSuccessfully(t *testing.T) {
	// Given
	repo := NewInMemoryCustomerRepository()
	newCustomer := domain.Customer{
		ID:     "new-id",
		Name:   "New Customer",
		Active: true,
	}

	// When
	err := repo.Save(newCustomer)
	savedCustomer, _ := repo.FindByID("new-id")

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, savedCustomer, "Expected a customer, got nil")
	assert.Equal(t, newCustomer.ID, savedCustomer.ID, "Expected customer ID to match")
	assert.Equal(t, newCustomer.Name, savedCustomer.Name, "Expected customer name to match")
	assert.Equal(t, newCustomer.Active, savedCustomer.Active, "Expected customer active status to match")
}

// TestSave_UpdateExistingCustomer_UpdatesSuccessfully
func TestSave_UpdateExistingCustomer_UpdatesSuccessfully(t *testing.T) {
	// Given
	repo := NewInMemoryCustomerRepository()
	existingCustomerID := "b7fc4cb6-6844-4cd0-95fb-f424a3938eb4"
	updatedCustomer := domain.Customer{
		ID:     existingCustomerID,
		Name:   "Updated Customer",
		Active: false,
	}

	// When
	err := repo.Save(updatedCustomer)
	savedCustomer, _ := repo.FindByID(existingCustomerID)

	// Then
	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, savedCustomer, "Expected a customer, got nil")
	assert.Equal(t, updatedCustomer.ID, savedCustomer.ID, "Expected customer ID to match")
	assert.Equal(t, updatedCustomer.Name, savedCustomer.Name, "Expected customer name to match")
	assert.Equal(t, updatedCustomer.Active, savedCustomer.Active, "Expected customer active status to match")
}
