package service

import (
	"errors"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateUser_Success tests the CreateUser service method when the Team exists
func TestCreateUser_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of TeamRepository to return a team
	mockTeamRepo.On("FindByID", uint(1)).Return(&entity.Team{
		ID:   1,
		Name: "Test Team",
	}, nil)

	// Mock the Create method of UserRepository to return nil (no error)
	mockUserRepo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

	// Create a new User entity
	user := &entity.User{
		Name:   "Test User",
		TeamID: 1,
	}

	// Call the CreateUser service method
	err := service.CreateUser(user)

	// Assert that there is no error
	assert.NoError(t, err)
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestCreateUser_TeamNotFound tests the CreateUser service method when the Team does not exist
func TestCreateUser_TeamNotFound(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of TeamRepository to return nil (team not found)
	mockTeamRepo.On("FindByID", uint(1)).Return(nil, nil)

	// Create a new User entity
	user := &entity.User{
		Name:   "Test User",
		TeamID: 1,
	}

	// Call the CreateUser service method
	err := service.CreateUser(user)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "team does not exist", err.Error())
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestCreateUser_TeamError tests the CreateUser service method when there is an error fetching the Team
func TestCreateUser_TeamError(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of TeamRepository to return an error
	mockTeamRepo.On("FindByID", uint(1)).Return(nil, errors.New("unable to find team"))

	// Create a new User entity
	user := &entity.User{
		Name:   "Test User",
		TeamID: 1,
	}

	// Call the CreateUser service method
	err := service.CreateUser(user)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Equal(t, "unable to find team", err.Error())
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByID tests the FindUserByID service method when the user is found
func TestFindUserByID_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of UserRepository to return a user
	mockUserRepo.On("FindByID", uint(1)).Return(&entity.User{
		ID:   1,
		Name: "Test User",
	}, nil)

	// Call the FindUserByID service method
	user, err := service.FindUserByID(1)

	// Assert that the user is found and there is no error
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, uint(1), user.ID)
	assert.Equal(t, "Test User", user.Name)
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByID_NotFound tests the FindUserByID service method when the user is not found
func TestFindUserByID_NotFound(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of UserRepository to return nil (user not found)
	mockUserRepo.On("FindByID", uint(1)).Return(nil, nil)

	// Call the FindUserByID service method
	user, err := service.FindUserByID(1)

	// Assert that the user is not found and there is no error
	assert.NoError(t, err)
	assert.Nil(t, user)
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByID_Error tests the FindUserByID service method when an error occurs
func TestFindUserByID_Error(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindByID method of UserRepository to return an error
	mockUserRepo.On("FindByID", uint(1)).Return(nil, errors.New("unable to find user"))

	// Call the FindUserByID service method
	user, err := service.FindUserByID(1)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, "unable to find user", err.Error())
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByTeamID tests the FindUserByTeamID service method
func TestFindUserByTeamID_Success(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindUserByTeamID method of UserRepository to return a list of users
	mockUserRepo.On("FindUserByTeamID", uint(1)).Return([]entity.User{
		{ID: 1, Name: "User 1", TeamID: 1},
		{ID: 2, Name: "User 2", TeamID: 1},
	}, nil)

	// Call the FindUserByTeamID service method
	users, err := service.FindUserByTeamID(1)

	// Assert that users are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "User 1", users[0].Name)
	assert.Equal(t, "User 2", users[1].Name)
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByTeamID_NoResults tests the FindUserByTeamID service method when no users are found
func TestFindUserByTeamID_NoResults(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindUserByTeamID method of UserRepository to return an empty list
	mockUserRepo.On("FindUserByTeamID", uint(1)).Return([]entity.User{}, nil)

	// Call the FindUserByTeamID service method
	users, err := service.FindUserByTeamID(1)

	// Assert that no users are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, users, 0)
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}

// TestFindUserByTeamID_Error tests the FindUserByTeamID service method when an error occurs
func TestFindUserByTeamID_Error(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	mockTeamRepo := new(mocks.TeamRepository)
	service := NewUserService(mockUserRepo, mockTeamRepo)

	// Mock the FindUserByTeamID method of UserRepository to return an error
	mockUserRepo.On("FindUserByTeamID", uint(1)).Return(nil, errors.New("unable to find users"))

	// Call the FindUserByTeamID service method
	users, err := service.FindUserByTeamID(1)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Equal(t, "unable to find users", err.Error())
	mockTeamRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
