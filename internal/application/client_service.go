package application

import (
	"customer-api/internal/domain"
	"customer-api/internal/ports"
)

type ClientService struct {
	repo ports.ClientRepository
}

func NewClientService(repo ports.ClientRepository) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) GetClient(id string) (*domain.Client, error) {
	return s.repo.GetClientByID(id)
}
