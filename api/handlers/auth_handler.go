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
	var password string = hashPassword(user.Password);

	log.Printf("User %s trying to login... with %s",username,password);

	//verify user


	//SQL injection
	query := fmt.Sprintf("SELECT userID FROM user WHERE username='%s' AND password='%s' ALLOW FILTERING",username,password);

	log.Println(query);

	iter := session.Query(query).WithContext(c).Iter();

	if iter.NumRows()==0 && !iter.Scan() {
		log.Println("ResultSet Empty")
		c.JSON(http.StatusBadRequest,gin.H{"status":"Authentication Failed"});
		return;

	}else if iter.NumRows()>1{
		//alert admin
		log.Println("ResultSet greater than 1")
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


func hashPassword(password string) string{
	return password;
}