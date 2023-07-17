package handlers

import (
	"log"
	"net/http"

	"api/models"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func GetData(c *gin.Context){

	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	query := "SELECT data FROM user WHERE username=? ALLOW FILTERING";

	resultSet := session.Query(query,username);

	var data []models.Data;

	err := resultSet.Scan(&data);
	
	if err!=nil{
		log.Println("Error while binding data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"GET data failed"})
		return;
	}

	c.JSON(http.StatusAccepted,data);
}