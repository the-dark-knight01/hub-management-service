package handler

import (
	"github.com/gin-gonic/gin"
	"hub_management_service/internal/entity"
	"hub_management_service/internal/service"
	"net/http"
	"strconv"
)

type TeamHandler struct {
	service service.TeamService
}

func NewTeamHandler(service service.TeamService) *TeamHandler {
	return &TeamHandler{service: service}
}

func (h *TeamHandler) CreateTeam(c *gin.Context) {
	var team entity.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateTeam(&team); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team created successfully", "team": team})
}

// FindTeamsByHubID - Endpoint to find teams by HubID
func (h *TeamHandler) FindTeamsByHubID(c *gin.Context) {
	hubID := c.Param("hub_id")
	if hubID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hub ID is required"})
		return
	}

	// Convert the hubID to uint
	hubIDUint, err := strconv.ParseUint(hubID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Hub ID"})
		return
	}

	teams, err := h.service.FindTeamsByHubID(uint(hubIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No teams found for this hub"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"teams": teams})
}

// FindTeamByID - Endpoint to find a team by its ID
func (h *TeamHandler) FindTeamByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Team ID is required"})
		return
	}

	teamIDUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Team ID"})
		return
	}

	team, err := h.service.FindByID(uint(teamIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if team == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"team": team})
}
