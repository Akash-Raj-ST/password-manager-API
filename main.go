package main

import (
	"log"
	"fmt"
	"os"

	"api/DB"
	"api/api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	session := DB.Connect();

	defer session.Close();

	router := gin.Default();
	routes.SetRoutes(router,session);

	// Start the server
	log.Println("API server running in",os.Getenv("API_PORT"))
	router.Run(fmt.Sprintf(":%s",os.Getenv("API_PORT")))
}