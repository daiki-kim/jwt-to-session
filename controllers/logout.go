package controllers

import (
	"awesomeProject/auth"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	err := auth.ClearSession(c)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"Message": "logout success"})
}
