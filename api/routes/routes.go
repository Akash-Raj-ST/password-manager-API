package routes

import (
	"log"

	"api/api/handlers"

	"github.com/gin-gonic/gin"
)

func StartAPI() {
	// Create a new Gin router
	router := gin.Default()

	// Define routes
	router.GET("/", handlers.AuthHandler)

	// Start the server
	log.Println("API server running in 8080")
	router.Run(":8080")
}
