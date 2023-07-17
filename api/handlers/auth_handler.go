package handlers

import (
	"api/api/models"
	"api/api/utils"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func AuthHandler(c *gin.Context) {
	//get the session

	session := c.MustGet("session").(*gocql.Session) 

	var user models.User;

	err := c.ShouldBindJSON(&user);

	if err!=nil {
		log.Println(err);

		c.JSON(http.StatusBadRequest,gin.H{"error":"Fields doesn't match"});
		return;
	}

	var username string = user.Username;
	password,err  := utils.GenerateHash(user.Password);

	if err!=nil {
		log.Println(err);

		c.JSON(http.StatusBadRequest,gin.H{"error":"password error"});
		return;
	}

	log.Printf("User %s trying to login... with %s",username,password);

	//verify user

	query := "SELECT password FROM user WHERE username=?";


	resultSet := session.Query(query,username);

	if resultSet.Iter().NumRows()==0 {
		log.Println("ResultSet Empty")
		c.JSON(http.StatusBadRequest,gin.H{"status":"Authentication Failed"});
		return;
	}

	var hashedPassword string;
	err = resultSet.Scan(&hashedPassword);
	if err!=nil{
		log.Println("Error while binding data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed"})
		return;
	}

	if hashedPassword!=password{
		c.JSON(http.StatusBadRequest,gin.H{"status":"Authentication Failed"});
		return;
	}

	// Generate a JWT
	tokenString, err := utils.GenerateJWT(username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated Token:", tokenString)

	c.JSON(http.StatusAccepted,gin.H{"session":tokenString});
}


