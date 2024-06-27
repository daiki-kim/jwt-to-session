package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"awesomeProject/auth"
)

func JwtTokenVerifier() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get("Authorization")

		bearerToken := strings.Split(authorizationHeader, "Bearer ")
		if len(bearerToken) != 2 {
			ctx.JSON(400, "invalid bearer token")
			ctx.Abort()
			return
		}
		token := strings.TrimSpace(bearerToken[1])

		claims, err := auth.ValidateToken(token, auth.TokenVerifyKey)
		if err != nil {
			ctx.JSON(400, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	}
}
