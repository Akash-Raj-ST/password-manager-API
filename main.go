package main

import (
	"api/DB"
	"api/api/routes"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	session := DB.Connect();

	routes.StartAPI(session)
}