package handlers

import (
	"api/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func CreateData(c *gin.Context){
	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	var newData models.Data;
	err := c.ShouldBindJSON(&newData);

	if err!=nil {
		log.Println(err);

		c.JSON(http.StatusBadRequest,gin.H{"error":"Fields doesn't match"});
		return;
	}

	query := "SELECT data FROM user WHERE username=?";

	resultSet := session.Query(query,username);

	var allData []models.Data;

	err = resultSet.Scan(&allData);

	if err!=nil{
		log.Println("Error while binding query data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed"})
		return;
	}

	newData.DataID = gocql.TimeUUID();
	allData = append(allData, newData);

	query = "UPDATE user SET data=? WHERE username=?;"

	err = session.Query(query,allData,username).Exec();
	if err!=nil{
		log.Println(err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed"})
		return;
	}

	c.JSON(http.StatusAccepted,allData);
}