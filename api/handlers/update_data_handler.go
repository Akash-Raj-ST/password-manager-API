package handlers

import (
	"fmt"
	"log"
	"net/http"

	"api/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)


func UpdateData(c *gin.Context){
	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	reqDataID := c.Param("data_id");
	if reqDataID==""{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"id required [/getData/:data_id]"})
		return;
	}
	
	var updateData models.Data;
	
	err := c.ShouldBindJSON(&updateData);
	if err!=nil{
		log.Println("Error while binding data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"Data didn't match"})
		return;
	}

	query := fmt.Sprintf("SELECT data FROM user WHERE username='%s' ALLOW FILTERING",username);

	resultSet := session.Query(query);

	var dataList []models.Data;

	err = resultSet.Scan(&dataList);
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Failed while binding"});
		return;
	}

	data_found := false;
	var updatedDataList []models.Data;

	for _,data := range dataList{

		if(data.DataID.String()==reqDataID){
			data_found = true;
			data = updateData;
		}
		updatedDataList = append(updatedDataList, data);
	}

	if !data_found{
		c.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Not such data exists"});
		return;
	}
	
	query = "UPDATE user SET data=? WHERE username=?";

	err = session.Query(query,updatedDataList,username).Exec();
	if err!=nil{
		log.Println("Error while updating data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"Update data failed"})
		return;
	}

	c.JSON(http.StatusAccepted,updatedDataList);
}