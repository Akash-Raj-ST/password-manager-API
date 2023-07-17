package handlers

import (
	"api/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func DeleteData(c *gin.Context){
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

	err := resultSet.Scan(&dataList);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Failed while binding"});
		return;
	}

	data_found := false;
	var newDataList []models.Data;

	for _,data := range dataList{

		if(data.DataID.String()==reqDataID){
			data_found = true;
			continue;
		}
		newDataList = append(newDataList, data);
	}

	if !data_found{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Not such data exists"});
		return;
	}


	query = "UPDATE user SET data=? WHERE username=?;";

	err = session.Query(query,newDataList,username).Exec();
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed"})
		return;
	}

	c.JSON(http.StatusAccepted,newDataList);
}