package application

import (
	"customer-api/internal/domain"
	"customer-api/internal/ports"
	"errors"
)

type CustomerServiceImpl struct {
	repo ports.CustomerRepository
}

func NewCustomerServiceImpl(repo ports.CustomerRepository) *CustomerServiceImpl {
	return &CustomerServiceImpl{repo: repo}
}

func (s *CustomerServiceImpl) CreateCustomer(customer domain.Customer) error {
	return s.repo.Save(customer)
}

func (s *CustomerServiceImpl) GetCustomer(id string) (*domain.Customer, error) {
	return s.repo.FindByID(id)
}

func (s *CustomerServiceImpl) ValidateCustomer(customerID string) error {
	if customerID == "" {
		return errors.New("ID de cliente inválido")
	}
	return nil
}
