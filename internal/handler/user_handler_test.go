package handler

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hub_management_service/internal/service/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateUser tests the CreateUser handler with valid input
func TestCreateUser(t *testing.T) {
	mockService := new(mocks.UserService)
	handler := NewUserHandler(mockService)

	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	// Mock the CreateUser behavior
	mockService.On("CreateUser", mock.AnythingOfType("*entity.User")).Return(nil)

	// Create request with valid JSON body
	body := `{
		"name": "Test User",
		"email": "testuser@example.com",
		"team_id": 1
	}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code and message
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "User created successfully")
	mockService.AssertExpectations(t)
}

// TestCreateUser_BadRequest tests the CreateUser handler with invalid input
func TestCreateUser_BadRequest(t *testing.T) {
	mockService := new(mocks.UserService)
	handler := NewUserHandler(mockService)

	router := gin.Default()
	router.POST("/users", handler.CreateUser)

	// Create invalid JSON body (missing email)
	body := `{
		"name": "Test User"
	}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(body))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert the response code for bad request
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	mockService.AssertExpectations(t)
}
