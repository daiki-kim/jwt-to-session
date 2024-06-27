package controllers

import (
	"awesomeProject/auth"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	err := auth.ClearSession(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Message": "logout success"})
}
