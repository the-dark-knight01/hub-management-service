package main

import (
	"hub_management_service/internal/handler"
	"hub_management_service/internal/repository"
	"hub_management_service/internal/router"
	"hub_management_service/internal/service"
	"hub_management_service/pkg/database"
	"log"
)

func main() {
	db := database.InitDB()
	defer database.CloseDB(db) // Ensure the DB is closed when the program ends

	hubRepo := repository.NewHubRepository(db)
	teamRepo := repository.NewTeamRepository(db)
	userRepo := repository.NewUserRepository(db)

	hubService := service.NewHubService(hubRepo)
	teamService := service.NewTeamService(teamRepo, hubRepo)
	userService := service.NewUserService(userRepo, teamRepo)

	hubHandler := handler.NewHubHandler(hubService)
	teamHandler := handler.NewTeamHandler(teamService)
	userHandler := handler.NewUserHandler(userService)

	r := router.NewRouter(hubHandler, teamHandler, userHandler)
	log.Fatal(r.Run(":8080"))
}
