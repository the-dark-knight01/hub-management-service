package service

import (
	"errors"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository"
)

type TeamService interface {
	CreateTeam(team *entity.Team) error
	FindTeamsByHubID(hubID uint) ([]entity.Team, error)
	FindByID(id uint) (*entity.Team, error)
}
type teamService struct {
	repo    repository.TeamRepository
	hubRepo repository.HubRepository // Add HubRepository to check Hub existence

}

func NewTeamService(repo repository.TeamRepository, hubRepo repository.HubRepository) TeamService {
	return &teamService{repo: repo, hubRepo: hubRepo}
}

func (s *teamService) CreateTeam(team *entity.Team) error {
	// Check if the Hub exists
	hub, err := s.hubRepo.FindByID(team.HubID)
	if err != nil || hub == nil {
		return errors.New("hub does not exist") // Return an error if the Hub does not exist
	}

	// Create the team if Hub exists
	return s.repo.Create(team)
}

func (s *teamService) FindTeamsByHubID(hubID uint) ([]entity.Team, error) {
	return s.repo.FindByHubID(hubID) // Call the repository method
}

func (s *teamService) FindByID(id uint) (*entity.Team, error) {
	return s.repo.FindByID(id) // Call the repository method
}
