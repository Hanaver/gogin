package router

import (
	"ggin/app/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	// 注册路由中间件
	r.Use(middleware.RouterMiddleware())
	
	// 注册api路由
	ApiRouter(r)

	// 注册web路由
	WebRouter(r)

	return r
}