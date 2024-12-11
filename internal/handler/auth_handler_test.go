package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Success(t *testing.T) {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		LoginHandler(c)
	})

	loginReq := map[string]string{"username": "admin", "password": "password"}
	reqBody, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]string
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NotEmpty(t, response["token"])

}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	router := gin.Default()

	router.POST("/login", func(c *gin.Context) {
		LoginHandler(c)
	})

	loginReq := map[string]string{"username": "wronguser", "password": "wrongpassword"}
	reqBody, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest("POST", "/login", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	var response map[string]string
	json.Unmarshal(resp.Body.Bytes(), &response)
	assert.Equal(t, "Invalid credentials", response["error"])
}
