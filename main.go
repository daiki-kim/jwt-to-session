package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"awesomeProject/auth"
	"awesomeProject/controllers"
	"awesomeProject/middlewares"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions(auth.SessionName, store))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1/")
		{
			v1.POST("/login", controllers.Login)
			v1.POST("/signup", controllers.Signup)
			v1.GET("/profile", middlewares.SessionAuth(), controllers.Profile)
			v1.POST("/logout", middlewares.SessionAuth(), controllers.Logout)
			// v1.GET("/profile", middlewares.JwtTokenVerifier(), controllers.Profile)
			// v1.GET("/profile", controllers.Profile)
		}
	}
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
