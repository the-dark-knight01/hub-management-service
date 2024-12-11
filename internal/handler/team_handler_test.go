package handler

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateTeam tests the CreateTeam handler with valid input
func TestCreateTeam(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.POST("/teams", handler.CreateTeam)

	// Mock the CreateTeam behavior
	mockService.On("CreateTeam", mock.AnythingOfType("*entity.Team")).Return(nil)

	// Create request with valid JSON body
	body := `{
		"name": "Test Team",
		"hub_id": 1
	}`
	req, _ := http.NewRequest("POST", "/teams", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and message
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Team created successfully")
	mockService.AssertExpectations(t)
}

// TestCreateTeam_BadRequest tests the CreateTeam handler with invalid input
func TestCreateTeam_BadRequest(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.POST("/teams", handler.CreateTeam)

	// Create invalid JSON body (missing hub_id)
	body := `{"name": "Test Team"}`
	req, _ := http.NewRequest("POST", "/teams", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for bad request
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertExpectations(t)
}

// TestFindTeamsByHubID tests the FindTeamsByHubID handler when teams are found
func TestFindTeamsByHubID(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.GET("/teams/:hub_id", handler.FindTeamsByHubID)

	// Mock the FindTeamsByHubID behavior
	mockService.On("FindTeamsByHubID", uint(1)).Return([]entity.Team{
		{
			ID:    1,
			Name:  "Test Team",
			HubID: 1,
		},
	}, nil)

	// Create request with HubID parameter
	req, _ := http.NewRequest("GET", "/teams/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and structure
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Team")
	mockService.AssertExpectations(t)
}

// TestFindTeamsByHubID_NotFound tests the FindTeamsByHubID handler when no teams are found
func TestFindTeamsByHubID_NotFound(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.GET("/teams/:hub_id", handler.FindTeamsByHubID)

	// Mock the FindTeamsByHubID behavior for not found case
	mockService.On("FindTeamsByHubID", uint(1)).Return([]entity.Team{}, nil)

	// Create request with HubID parameter
	req, _ := http.NewRequest("GET", "/teams/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for not found
	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

// TestFindTeamByID tests the FindTeamByID handler when the team is found
func TestFindTeamByID(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.GET("/teams/:id", handler.FindTeamByID)

	// Mock the FindTeamByID behavior
	mockService.On("FindByID", uint(1)).Return(&entity.Team{
		ID:    1,
		Name:  "Test Team",
		HubID: 1,
	}, nil)

	// Create request with ID parameter
	req, _ := http.NewRequest("GET", "/teams/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and structure
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Team")
	mockService.AssertExpectations(t)
}

// TestFindTeamByID_NotFound tests the FindTeamByID handler when the team is not found
func TestFindTeamByID_NotFound(t *testing.T) {
	mockService := new(mocks.TeamService)
	handler := NewTeamHandler(mockService)

	router := gin.Default()
	router.GET("/teams/:id", handler.FindTeamByID)

	// Mock the FindTeamByID behavior for not found case
	mockService.On("FindByID", uint(1)).Return(nil, nil)

	// Create request with ID parameter
	req, _ := http.NewRequest("GET", "/teams/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for not found
	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}
