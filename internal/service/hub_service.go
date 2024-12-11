package service

import (
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository"
)

type HubService interface {
	CreateHub(hub *entity.Hub) error
	FindHubByID(id uint) (*entity.Hub, error)
	SearchHubsByName(name string) ([]entity.Hub, error)
}

type hubService struct {
	repo repository.HubRepository
}

func NewHubService(repo repository.HubRepository) HubService {
	return &hubService{repo: repo}
}

func (s *hubService) CreateHub(hub *entity.Hub) error {
	return s.repo.Create(hub)
}

// FindHubByID fetches a hub by its ID
func (s *hubService) FindHubByID(id uint) (*entity.Hub, error) {
	return s.repo.FindByID(id)
}

// SearchHubsByName searches for hubs by name
func (s *hubService) SearchHubsByName(name string) ([]entity.Hub, error) {
	return s.repo.SearchByName(name)
}
