package ports

import "customer-api/internal/domain"

//go:generate mockery --name=ClientRepository --output=. --outpkg=ports --filename=mock_customer_repository.go
type ClientRepository interface {
	GetClientByID(id string) (*domain.Client, error)
}
