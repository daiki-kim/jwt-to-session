package controllers

import (
	"github.com/gin-gonic/gin"

	"awesomeProject/models"
)

func Signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	err = user.CreateUser()
	if err != nil {
		ctx.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"Message": "signup success"})
}
