package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/channyeintun/golang-jwt-authentication/config"
	"github.com/channyeintun/golang-jwt-authentication/db"
	"github.com/channyeintun/golang-jwt-authentication/handlers"
	"github.com/channyeintun/golang-jwt-authentication/repositories"
	"github.com/channyeintun/golang-jwt-authentication/services"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "JWT Authentication",
		ServerHeader: "Fiber",
	})

	// Repositories
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	// privateRoutes := server.Use(middlewares.AuthProtected(db))

	// handlers.NewEventHandler(privateRoutes.Group("/protected"), someRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
