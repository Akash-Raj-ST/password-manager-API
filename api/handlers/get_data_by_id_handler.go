package handlers

import (
	"api/models"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func GetDataByID(c *gin.Context){
	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	reqDataID := c.Param("data_id");

	if reqDataID==""{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"id required [/getData/:data_id]"})
		return;
	}

	query := "SELECT data FROM user WHERE username=? ALLOW FILTERING";

	resultSet := session.Query(query,username);

	var dataList []models.Data;
	var reqData models.Data;

	err := resultSet.Scan(&dataList);
	if err!=nil{
		log.Println("Error while retreiving data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Failed while binding"});
		return;
	}

	data_found := false;

	for _,data := range dataList{

		if(data.DataID.String()==reqDataID){
			data_found = true;
			reqData = data;
			break;
		}

		log.Println(data.DataID.String());
	
	}

	if !data_found{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Not such data exists"});
			return;
	}

	c.JSON(http.StatusAccepted,reqData);
}