package router

import (
	"ggin/app/handler"
	"ggin/app/middleware"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	
    api.GET("/subexample", handler.SubExampleHandler)
	api.GET("/home", handler.HomeHandler)
	api.POST("/register", handler.RegisterHandler)
	api.POST("/login", handler.LoginHandler)

	auth := api.Group("/auth")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", handler.ProfileHandler)
		// auth.POST("/update", handler.UpdateProfileHandler)
	}
}
