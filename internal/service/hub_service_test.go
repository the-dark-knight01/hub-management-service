package service

import (
	"errors"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateHub tests the CreateHub service method
func TestCreateHub(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the Create method of HubRepository
	mockRepo.On("Create", mock.AnythingOfType("*entity.Hub")).Return(nil)

	// Create a new Hub entity
	hub := &entity.Hub{
		Name: "Test Hub",
	}

	// Call the CreateHub service method
	err := service.CreateHub(hub)

	// Assert that there is no error
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestCreateHub_Error tests the CreateHub service method when the repository returns an error
func TestCreateHub_Error(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the Create method of HubRepository to return an error
	mockRepo.On("Create", mock.AnythingOfType("*entity.Hub")).Return(errors.New("unable to create hub"))

	// Create a new Hub entity
	hub := &entity.Hub{
		Name: "Test Hub",
	}

	// Call the CreateHub service method
	err := service.CreateHub(hub)

	// Assert that the error is returned
	assert.Error(t, err)
	assert.Equal(t, "unable to create hub", err.Error())
	mockRepo.AssertExpectations(t)
}

// TestFindHubByID tests the FindHubByID service method when the hub is found
func TestFindHubByID(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the FindByID method of HubRepository to return a hub
	mockRepo.On("FindByID", uint(1)).Return(&entity.Hub{
		ID:   1,
		Name: "Test Hub",
	}, nil)

	// Call the FindHubByID service method
	hub, err := service.FindHubByID(1)

	// Assert that the hub is found and there is no error
	assert.NoError(t, err)
	assert.NotNil(t, hub)
	assert.Equal(t, uint(1), hub.ID)
	assert.Equal(t, "Test Hub", hub.Name)
	mockRepo.AssertExpectations(t)
}

// TestFindHubByID_NotFound tests the FindHubByID service method when the hub is not found
func TestFindHubByID_NotFound(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the FindByID method of HubRepository to return nil (hub not found)
	mockRepo.On("FindByID", uint(1)).Return(nil, nil)

	// Call the FindHubByID service method
	hub, err := service.FindHubByID(1)

	// Assert that the hub is not found and there is no error
	assert.NoError(t, err)
	assert.Nil(t, hub)
	mockRepo.AssertExpectations(t)
}

// TestFindHubByID_Error tests the FindHubByID service method when an error occurs
func TestFindHubByID_Error(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the FindByID method of HubRepository to return an error
	mockRepo.On("FindByID", uint(1)).Return(nil, errors.New("unable to find hub"))

	// Call the FindHubByID service method
	hub, err := service.FindHubByID(1)

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, hub)
	assert.Equal(t, "unable to find hub", err.Error())
	mockRepo.AssertExpectations(t)
}

// TestSearchHubsByName tests the SearchHubsByName service method when hubs are found
func TestSearchHubsByName(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the SearchByName method of HubRepository to return a list of hubs
	mockRepo.On("SearchByName", "Test").Return([]entity.Hub{
		{ID: 1, Name: "Test Hub 1"},
		{ID: 2, Name: "Test Hub 2"},
	}, nil)

	// Call the SearchHubsByName service method
	hubs, err := service.SearchHubsByName("Test")

	// Assert that hubs are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, hubs, 2)
	assert.Equal(t, "Test Hub 1", hubs[0].Name)
	assert.Equal(t, "Test Hub 2", hubs[1].Name)
	mockRepo.AssertExpectations(t)
}

// TestSearchHubsByName_NoResults tests the SearchHubsByName service method when no hubs are found
func TestSearchHubsByName_NoResults(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the SearchByName method of HubRepository to return an empty list
	mockRepo.On("SearchByName", "NonExistent").Return([]entity.Hub{}, nil)

	// Call the SearchHubsByName service method
	hubs, err := service.SearchHubsByName("NonExistent")

	// Assert that no hubs are returned and there is no error
	assert.NoError(t, err)
	assert.Len(t, hubs, 0)
	mockRepo.AssertExpectations(t)
}

// TestSearchHubsByName_Error tests the SearchHubsByName service method when an error occurs
func TestSearchHubsByName_Error(t *testing.T) {
	mockRepo := new(mocks.HubRepository)
	service := NewHubService(mockRepo)

	// Mock the SearchByName method of HubRepository to return an error
	mockRepo.On("SearchByName", "Test").Return(nil, errors.New("unable to search hubs"))

	// Call the SearchHubsByName service method
	hubs, err := service.SearchHubsByName("Test")

	// Assert that an error is returned
	assert.Error(t, err)
	assert.Nil(t, hubs)
	assert.Equal(t, "unable to search hubs", err.Error())
	mockRepo.AssertExpectations(t)
}
