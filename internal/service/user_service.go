package service

import (
	"errors"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository"
)

type UserService interface {
	CreateUser(user *entity.User) error
	FindUserByID(id uint) (*entity.User, error)
	FindUserByTeamID(teamID uint) ([]entity.User, error)
}
type userService struct {
	repo     repository.UserRepository
	teamRepo repository.TeamRepository // Add team repository to check if a team exists
}

func NewUserService(repo repository.UserRepository, teamRepo repository.TeamRepository) UserService {
	return &userService{repo: repo, teamRepo: teamRepo}
}

func (s *userService) CreateUser(user *entity.User) error {
	// Check if the team exists before creating the user
	team, err := s.teamRepo.FindByID(user.TeamID)
	if err != nil {
		return err
	}
	if team == nil {
		return errors.New("team does not exist")
	}

	// Proceed to create the user if the team exists
	return s.repo.Create(user)
}

func (s *userService) FindUserByID(id uint) (*entity.User, error) {
	return s.repo.FindByID(id)
}

// FindUserByTeamID New method to find users by TeamID
func (s *userService) FindUserByTeamID(teamID uint) ([]entity.User, error) {
	// Find users by TeamID using the repository
	return s.repo.FindUserByTeamID(teamID)
}
