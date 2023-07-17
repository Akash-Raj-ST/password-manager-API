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

	if len(os.Args) >1 {
        input := os.Args[1]
		if(input=="initDB"){
			DB.InitializeDB();
		}else{
			log.Fatal("No such command Available:['initDB']")
		}
    }

	session := DB.ConnectDB();

	defer session.Close();

	router := gin.Default();
	routes.SetRoutes(router,session);

	// Start the server
	log.Println("API server running in",os.Getenv("API_PORT"))
	router.Run(fmt.Sprintf(":%s",os.Getenv("API_PORT")))
}