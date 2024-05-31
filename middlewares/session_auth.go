package middlewares

import (
	"awesomeProject/auth"

	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, err := auth.GetSession(c)
		if err != nil {
			c.JSON(401, gin.H{"Error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("email", email)
		c.Next()
	}
}
