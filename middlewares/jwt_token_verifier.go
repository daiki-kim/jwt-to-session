package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"awesomeProject/auth"
)

func JwtTokenVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		bearerToken := strings.Split(authorizationHeader, "Bearer ")
		if len(bearerToken) != 2 {
			c.JSON(400, "invalid bearer token")
			c.Abort()
			return
		}
		token := strings.TrimSpace(bearerToken[1])

		claims, err := auth.ValidateToken(token, auth.TokenVerifyKey)
		if err != nil {
			c.JSON(400, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
