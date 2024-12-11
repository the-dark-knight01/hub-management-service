package handler

import (
	"github.com/gin-gonic/gin"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/service"
	"net/http"
	"strconv"
)

type HubHandler struct {
	service service.HubService
}

func NewHubHandler(service service.HubService) *HubHandler {
	return &HubHandler{service: service}
}

// CreateHub handles the creation of a new hub
func (h *HubHandler) CreateHub(c *gin.Context) {
	var hub entity.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		// If binding fails, return 400 with the error message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateHub(&hub); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Hub created successfully", "hub": hub})
}

// FindHubByID retrieves a hub by its ID
func (h *HubHandler) FindHubByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	hub, err := h.service.FindHubByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if hub == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Hub not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hub": hub})
}

// SearchHubsByName searches for hubs by name
func (h *HubHandler) SearchHubsByName(c *gin.Context) {
	name := c.DefaultQuery("name", "") // Get the 'name' query parameter
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required"})
		return
	}

	hubs, err := h.service.SearchHubsByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(hubs) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hubs found with the given name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"hubs": hubs})
}
