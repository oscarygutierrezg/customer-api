package ports

import "customer-api/internal/domain"

//go:generate mockery --name=CustomerRepository --output=. --outpkg=ports --filename=mock_customer_repository.go
type CustomerRepository interface {
	Save(customer domain.Customer) error
	FindByID(id string) (*domain.Customer, error)
}
