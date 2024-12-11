package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hub_management_service/internal/handler"
	"hub_management_service/internal/middleware"
)

// NewRouter initializes and returns the Gin router with all routes and middleware applied
func NewRouter(hubHandler *handler.HubHandler, teamHandler *handler.TeamHandler, userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()
	// Custom CORS configuration using gin-contrib/cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8081"}                   // Allow swagger UI
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"} // Allow necessary methods
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}           // Allow the Authorization header
	corsConfig.ExposeHeaders = []string{"Authorization"}                          // Expose the Authorization header if needed

	// Apply CORS middleware to the Gin router
	r.Use(cors.New(corsConfig))
	// Login route (no auth required)
	r.POST("/login", handler.LoginHandler)

	// Protected routes with authentication middleware
	r.POST("/hubs", middleware.AuthMiddleware(), hubHandler.CreateHub)
	r.GET("/hubs/:id", hubHandler.FindHubByID)         // Get hub by ID
	r.GET("/hubs/search", hubHandler.SearchHubsByName) // Search hubs by name

	r.POST("/teams", middleware.AuthMiddleware(), teamHandler.CreateTeam)
	r.GET("/teams/hub/:hub_id", teamHandler.FindTeamsByHubID) // Find teams by hub ID
	r.GET("/teams/:id", teamHandler.FindTeamByID)             // Find team by ID

	r.POST("/users", middleware.AuthMiddleware(), userHandler.CreateUser)
	r.GET("/users/team/:team_id", userHandler.FindUserByTeamID) // Find users by team ID
	r.GET("/users/:id", userHandler.FindUserByID)               // Get user by ID

	return r
}
