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
	router.POST("/createData", handlers.CreateData)
	router.GET("/getData", handlers.GetData);
	router.GET("/getDataByID/:data_id",handlers.GetDataByID)
	router.POST("/updateData/:data_id", handlers.UpdateData)
	router.POST("/deleteData/:data_id", handlers.DeleteData)
	
}
