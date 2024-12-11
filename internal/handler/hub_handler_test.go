package handler

import (
	"bytes"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestFindHubByID tests the FindHubByID handler when the hub is found
func TestFindHubByID(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.GET("/hubs/:id", handler.FindHubByID)

	// Mock the FindHubByID behavior
	mockService.On("FindHubByID", uint(1)).Return(&entity.Hub{
		ID:       1,
		Name:     "Test Hub",
		Location: "Test Location",
	}, nil)

	// Create request with ID parameter
	req, _ := http.NewRequest("GET", "/hubs/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and structure
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Hub")
	mockService.AssertExpectations(t)
}

// TestFindHubByID_NotFound tests the FindHubByID handler when no hub is found
func TestFindHubByID_NotFound(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.GET("/hubs/:id", handler.FindHubByID)

	// Mock the FindHubByID behavior for not found case
	mockService.On("FindHubByID", uint(1)).Return(nil, nil)

	// Create request with ID parameter
	req, _ := http.NewRequest("GET", "/hubs/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for not found
	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

// TestSearchHubsByName tests the SearchHubsByName handler when hubs are found
func TestSearchHubsByName(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.GET("/hubs/search", handler.SearchHubsByName)

	// Mock the SearchHubsByName behavior
	mockService.On("SearchHubsByName", "Test Hub").Return([]entity.Hub{
		{
			ID:       1,
			Name:     "Test Hub",
			Location: "Test Location",
		},
	}, nil)

	// Create request with query parameter
	req, _ := http.NewRequest("GET", "/hubs/search?name=Test Hub", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and structure
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Test Hub")
	mockService.AssertExpectations(t)
}

// TestSearchHubsByName_NoResults tests the SearchHubsByName handler when no hubs are found
func TestSearchHubsByName_NoResults(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.GET("/hubs/search", handler.SearchHubsByName)

	// Mock the SearchHubsByName behavior for no results
	mockService.On("SearchHubsByName", "Nonexistent Hub").Return([]entity.Hub{}, nil)

	// Create request with query parameter
	req, _ := http.NewRequest("GET", "/hubs/search?name=Nonexistent Hub", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for no results
	assert.Equal(t, http.StatusNotFound, resp.Code)
	mockService.AssertExpectations(t)
}

// TestCreateHub tests the CreateHub handler with valid input
func TestCreateHub(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.POST("/hubs", handler.CreateHub)

	// Mock the CreateHub behavior
	mockService.On("CreateHub", &entity.Hub{
		Name:     "Test Hub",
		Location: "Test Location",
	}).Return(nil)

	// Create request with valid JSON body
	body := `{"name": "Test Hub", "location": "Test Location"}`
	req, _ := http.NewRequest("POST", "/hubs", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and message
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Hub created successfully")
	mockService.AssertExpectations(t)
}

// TestCreateHub_BadRequest tests the CreateHub handler with invalid input
func TestCreateHub_BadRequest(t *testing.T) {
	mockService := new(mocks.HubService)
	handler := NewHubHandler(mockService)

	router := gin.Default()
	router.POST("/hubs", handler.CreateHub)

	// Create invalid JSON body (missing location field)
	body := `{"name": "Test Hub"}`
	req, _ := http.NewRequest("POST", "/hubs", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for bad request
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertExpectations(t)
}
