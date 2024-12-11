package service

import (
	"errors"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateTeam tests the CreateTeam service method when the Hub exists
func TestCreateTeam_Success(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of HubRepository to return a hub
	mockHubRepo.On("FindByID", uint(1)).Return(&entity.Hub{
		ID:   1,
		Name: "Test Hub",
	}, nil)

	// Mock the Create method of TeamRepository to return nil (no error)
	mockTeamRepo.On("Create", mock.AnythingOfType("*entity.Team")).Return(nil)

	// Create a new Team entity
	team := &entity.Team{
		Name:  "Test Team",
		HubID: 1,
	}

	// Call the CreateTeam service method
	err := service.CreateTeam(team)

	// Assert that there is no error
	assert.NoError(t, err)
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestCreateTeam_HubNotFound tests the CreateTeam service method when the Hub does not exist
func TestCreateTeam_HubNotFound(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of HubRepository to return nil (hub not found)
	mockHubRepo.On("FindByID", uint(1)).Return(nil, nil)

	// Create a new Team entity
	team := &entity.Team{
		Name:  "Test Team",
		HubID: 1,
	}

	// Call the CreateTeam service method
	err := service.CreateTeam(team)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "hub does not exist", err.Error())
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestCreateTeam_HubError tests the CreateTeam service method when there is an error fetching the Hub
func TestCreateTeam_HubError(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of HubRepository to return an error
	mockHubRepo.On("FindByID", uint(1)).Return(nil, errors.New("hub does not exist"))

	// Create a new Team entity
	team := &entity.Team{
		Name:  "Test Team",
		HubID: 1,
	}

	// Call the CreateTeam service method
	err := service.CreateTeam(team)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "hub does not exist", err.Error())
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindTeamsByHubID tests the FindTeamsByHubID service method
func TestFindTeamsByHubID(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByHubID method of TeamRepository to return a list of teams
	mockTeamRepo.On("FindByHubID", uint(1)).Return([]entity.Team{
		{ID: 1, Name: "Team 1", HubID: 1},
		{ID: 2, Name: "Team 2", HubID: 1},
	}, nil)

	// Call the FindTeamsByHubID service method
	teams, err := service.FindTeamsByHubID(1)

	// Assert that teams are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, teams, 2)
	assert.Equal(t, "Team 1", teams[0].Name)
	assert.Equal(t, "Team 2", teams[1].Name)
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindTeamsByHubID_NoResults tests the FindTeamsByHubID service method when no teams are found
func TestFindTeamsByHubID_NoResults(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByHubID method of TeamRepository to return an empty list
	mockTeamRepo.On("FindByHubID", uint(1)).Return([]entity.Team{}, nil)

	// Call the FindTeamsByHubID service method
	teams, err := service.FindTeamsByHubID(1)

	// Assert that no teams are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, teams, 0)
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindTeamsByHubID_Error tests the FindTeamsByHubID service method when an error occurs
func TestFindTeamsByHubID_Error(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByHubID method of TeamRepository to return an error
	mockTeamRepo.On("FindByHubID", uint(1)).Return(nil, errors.New("unable to find teams"))

	// Call the FindTeamsByHubID service method
	teams, err := service.FindTeamsByHubID(1)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, teams)
	assert.Equal(t, "unable to find teams", err.Error())
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindByID tests the FindByID service method when the team is found
func TestFindByID_Success(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of TeamRepository to return a team
	mockTeamRepo.On("FindByID", uint(1)).Return(&entity.Team{
		ID:   1,
		Name: "Test Team",
	}, nil)

	// Call the FindByID service method
	team, err := service.FindByID(1)

	// Assert that the team is found and there is no error
	assert.NoError(t, err)
	assert.NotNil(t, team)
	assert.Equal(t, uint(1), team.ID)
	assert.Equal(t, "Test Team", team.Name)
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindByID_NotFound tests the FindByID service method when the team is not found
func TestFindByID_NotFound(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of TeamRepository to return nil (team not found)
	mockTeamRepo.On("FindByID", uint(1)).Return(nil, nil)

	// Call the FindByID service method
	team, err := service.FindByID(1)

	// Assert that the team is not found and there is no error
	assert.NoError(t, err)
	assert.Nil(t, team)
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}

// TestFindByID_Error tests the FindByID service method when an error occurs
func TestFindByID_Error(t *testing.T) {
	mockTeamRepo := new(mocks.TeamRepository)
	mockHubRepo := new(mocks.HubRepository)
	service := NewTeamService(mockTeamRepo, mockHubRepo)

	// Mock the FindByID method of TeamRepository to return an error
	mockTeamRepo.On("FindByID", uint(1)).Return(nil, errors.New("unable to find team"))

	// Call the FindByID service method
	team, err := service.FindByID(1)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, team)
	assert.Equal(t, "unable to find team", err.Error())
	mockHubRepo.AssertExpectations(t)
	mockTeamRepo.AssertExpectations(t)
}
