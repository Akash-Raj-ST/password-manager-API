package handlers

import (
	"fmt"
	"log"
	"net/http"

	"api/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func GetData(c *gin.Context){

	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	query := fmt.Sprintf("SELECT data FROM user WHERE username='%s' ALLOW FILTERING",username);

	resultSet := session.Query(query);

	var data models.Data;

	err := resultSet.Scan(&data);

	if err!=nil{
		log.Println("Error while retreiving data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"GET data failed"})
		return;
	}

	c.JSON(http.StatusAccepted,data);
}