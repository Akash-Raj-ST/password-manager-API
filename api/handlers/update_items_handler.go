package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"api/api/models"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)


func UpdateData(c *gin.Context){
	session := c.MustGet("session").(*gocql.Session) 
	username := c.MustGet("username");

	
	var data models.Data;
	
	err := c.ShouldBindJSON(&data);


	if err!=nil{
		log.Println("Error while binding data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"Data didn't match"})
		return;
	}

	var itemStrings []string
	for _, item := range data.Items {
		itemStrings = append(itemStrings, fmt.Sprintf("{Key:'%s', Value:'%s' ,is_secret:%v}", item.Key, item.Value, item.IsSecret))
	}
	
	query := fmt.Sprintf("UPDATE user SET data={Username:'%s' , Password:'%s' , items:[%s]} WHERE username='%s';",
	data.Username, data.Password, strings.Join(itemStrings, ", "), username)

	log.Println(query);

	err = session.Query(query).Exec();

	if err!=nil{
		log.Println("Error while updating data",err.Error());
		c.JSON(http.StatusBadRequest,gin.H{"status":"Update data failed"})
		return;
	}

	c.JSON(http.StatusAccepted,data);
}