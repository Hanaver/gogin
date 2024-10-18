package handler

import (
	"github.com/gin-gonic/gin"
)

// SubExampleHandler 子路由示例处理函数
func SubExampleHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "subexample",
    })
}