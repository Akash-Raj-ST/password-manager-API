package handlers


import (
	"github.com/gin-gonic/gin"
)

func AuthHandler(c *gin.Context) {
	c.String(200, "Welcome to the home page!")
}