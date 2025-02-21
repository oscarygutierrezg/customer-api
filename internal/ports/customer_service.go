package ports

import "customer-api/internal/domain"

//go:generate mockery --name=CustomerService --output=. --outpkg=ports --filename=mock_customer_service.go
type CustomerService interface {
	CreateCustomer(customer domain.Customer) error
	GetCustomer(id string) (*domain.Customer, error)
	ValidateCustomer(customerID string) error
}
