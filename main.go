// main.go
package main

import (
	"friends-management-api/config"
	router "friends-management-api/modules"
	"friends-management-api/modules/auth"
	"friends-management-api/modules/auth/auth_handler"
	"friends-management-api/modules/auth/auth_repository"
	"friends-management-api/modules/auth/auth_service"
	"friends-management-api/modules/friend"
	"friends-management-api/modules/friend/friend_handler"
	"friends-management-api/modules/friend/friend_repository"
	"friends-management-api/modules/friend/friend_service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	r := gin.Default()

	authRepo := auth_repository.New(db)
	authService := auth_service.New(authRepo)
	authHandler := auth_handler.New(authService)
	authRoutes := auth.New(authHandler)

	friendRepo := friend_repository.New(db)
	friendService := friend_service.New(friendRepo, authRepo)
	friendHandler := friend_handler.New(friendService)
	friendRoutes := friend.New(friendHandler)

	// Initialize all routes and register them
	routes := router.NewRoutes(authRoutes, friendRoutes)
	routes.Setup(r)

	// Start the server
	r.Run()
}
