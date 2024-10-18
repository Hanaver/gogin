package router

import (
	"ggin/app/handler"

	"github.com/gin-gonic/gin"
)

func WebRouter(r *gin.Engine) {
	
	r.GET("/example", handler.ExampleHandler)
}