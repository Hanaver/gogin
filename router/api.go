package router

import (
	"ggin/app/handler"

	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	api := r.Group("/api")
	
    api.GET("/subexample", handler.SubExampleHandler)
	api.GET("/home", handler.HomeHandler)
	api.POST("/register", handler.RegisterHandler)
}