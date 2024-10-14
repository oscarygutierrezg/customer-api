package repository

import (
	"customer-api/internal/domain"
	"errors"
	"math/rand"
	"time"
)

type InMemoryClientRepository struct {
	customers map[string]*domain.Client
}

var names = []string{
	"Jose", "Alisson", "Luis", "Evelin", "Oscar", "Carmen",
}

var customerIds = []string{
	"b7fc4cb6-6844-4cd0-95fb-f424a3938eb4",
	"7ac66acd-411d-4d4b-9cc9-4f9cbe5425f7",
	"22a2cfa8-f577-448b-aa4a-f85a82c38656",
	"99d3c265-b0e2-4a27-8e1e-36fb58936975",
	"7cb8de15-2f66-4e71-9dc4-d86d1ff75364",
	"5b385b21-895c-4132-942c-79856eb9d287",
}

func NewInMemoryClientRepository() *InMemoryClientRepository {
	rand.Seed(time.Now().UnixNano())
	customers := make(map[string]*domain.Client)

	for i := 0; i < 6; i++ {
		customers[customerIds[i]] = &domain.Client{
			ID:     customerIds[i],
			Name:   randomName(),
			Active: i%2 == 0,
		}
	}

	return &InMemoryClientRepository{
		customers: customers,
	}
}

func randomName() string {
	return names[rand.Intn(len(names))]
}

func (r *InMemoryClientRepository) GetClientByID(id string) (*domain.Client, error) {
	customer, exists := r.customers[id]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}
