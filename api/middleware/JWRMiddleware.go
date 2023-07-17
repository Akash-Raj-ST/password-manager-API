package middleware

import (
	"api/utils"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)


func JWTMiddleware(session *gocql.Session) gin.HandlerFunc {

	authNotRequiredArray := [1]string {"/auth"}

    return func(c *gin.Context) {

		for _,path:=range(authNotRequiredArray){
			if path == c.FullPath(){
				c.Next();
				return;
			}
		}

		log.Println("JWT handling...");

		tokens:= c.Request.Header["Token"];

		if len(tokens)==0 {
			log.Println("Token missing");

			c.JSON(http.StatusBadRequest,gin.H{"status":"Token Authentication Failed"});
			return;
		}

		// Validate and decode the JWT
		claims, err := utils.ValidateJWT(tokens[0])

		if err != nil {
			log.Println(err.Error());

			c.JSON(http.StatusBadRequest,gin.H{"status":"Token Authentication Failed"});
			return;
		}

		c.Set("username", claims["username"]);

		c.Next();
    }
}