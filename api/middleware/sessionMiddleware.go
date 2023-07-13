package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func SessionMiddleware(session *gocql.Session) gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Println("session set");
        c.Set("session", session)
        c.Next()
    }
}