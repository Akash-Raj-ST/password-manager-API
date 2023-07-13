package routes

import (
	"log"
	"os"
	"fmt"

	"api/api/handlers"
	"api/api/middleware"

	"github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
)

func StartAPI(s *gocql.Session) {
	// Create a new Gin router
	router := gin.Default()


	router.Use(middleware.SessionMiddleware(s));
	router.Use(middleware.JWTMiddleware(s));
	// Define routes
	router.GET("/auth", handlers.AuthHandler)

	// Start the server
	log.Println("API server running in 8080")
	router.Run(fmt.Sprintf(":%s",os.Getenv("API_PORT")))
}
