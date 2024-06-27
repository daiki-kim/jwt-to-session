package controllers

import (
	"github.com/gin-gonic/gin"

	"awesomeProject/models"
)

func Profile(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	user, err := models.GetUserByEmail(email.(string))
	if err != nil {
		ctx.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	ctx.JSON(200, user)
}
