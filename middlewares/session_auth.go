package middlewares

import (
	"awesomeProject/auth"

	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, err := auth.GetSession(ctx)
		if err != nil {
			ctx.JSON(401, gin.H{"Error": "unauthorized"})
			ctx.Abort()
			return
		}

		if session == "" {
			ctx.JSON(401, gin.H{"Error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("email", session)

		ctx.Next()
	}
}
