package routes

import (
	"api/api/handlers"
	"api/api/middleware"

	"github.com/gocql/gocql"
	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine,s *gocql.Session) {

	router.Use(middleware.SessionMiddleware(s));
	router.Use(middleware.JWTMiddleware(s));
	
	// Define routes
	router.POST("/auth", handlers.AuthHandler);
	router.GET("/getData", handlers.GetData);
	router.POST("/updateData", handlers.UpdateData)
	
}
