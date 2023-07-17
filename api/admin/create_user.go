package main

import (
	"api/DB"
	"api/utils"

	"fmt"
	"log"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	session := DB.ConnectDB()

	defer session.Close()

	var username string
	var password string

	fmt.Print("Enter username: ")
	_,err := fmt.Scan(&username)
	if err != nil {
		log.Fatal(err.Error())
	}
	if username == "" {
		log.Fatal("username cannot be Empty!!!")
	}

	fmt.Print("Enter password: ")
	_,err = fmt.Scan(&password)
	if err != nil {
		log.Fatal(err.Error())
	}
	if password == "" {
		log.Fatal("Password cannot be Empty!!!")
	}

	password, err = utils.GenerateHash(password)
	if err != nil {
		log.Fatal("Error while Generating Hash", err.Error())
	}

	query := "INSERT INTO user(userid,username,password) VALUES(?,?,?)"
	err = session.Query(query,gocql.TimeUUID(),username, password).Exec()

	if err != nil {
		log.Fatal("Error while creating user", err.Error())
	}

	log.Println("User created successfully")
}
