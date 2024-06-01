package middlewares

import (
	"awesomeProject/auth"

	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := auth.GetSession(c)
		if err != nil {
			c.JSON(401, gin.H{"Error": "unauthorized"})
			c.Abort()
			return
		}

		if session == "" {
			c.JSON(401, gin.H{"Error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("email", session)

		c.Next()
	}
}
