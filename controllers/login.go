package controllers

import (
	"github.com/gin-gonic/gin"

	"awesomeProject/auth"
	"awesomeProject/models"
)

func Login(c *gin.Context) {
	loginRequestBody := &struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}
	err := c.ShouldBindJSON(&loginRequestBody)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	user, err := models.GetUserByEmail(loginRequestBody.Email)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	err = user.CompareHashAndPassword(loginRequestBody.Password)
	if err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	err = auth.SetSession(c, user.Email)
	if err != nil {
		c.JSON(500, gin.H{"Error": "unable to set session"})
		return
	}
	c.JSON(200, gin.H{"Message": "login success"})
}

// 	claim := auth.NewClaim(user.Email)
// 	token, err := claim.GenerateToken()
// 	if err != nil {
// 		c.JSON(500, gin.H{"Error": err.Error()})
// 		return
// 	}
// 	refreshCustomClaim := auth.NewClaim(user.Email)
// 	refreshToken, err := refreshCustomClaim.GenerateRefreshToken()
// 	if err != nil {
// 		c.JSON(500, gin.H{"Error": err.Error()})
// 		return
// 	}

// 	c.JSON(200, &struct {
// 		Token        string `json:"token"`
// 		RefreshToken string `json:"refreshToken"`
// 	}{
// 		Token:        token,
// 		RefreshToken: refreshToken,
// 	})
// }
