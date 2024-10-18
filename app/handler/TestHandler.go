package handler

import (
	"ggin/app/utils"

	"github.com/gin-gonic/gin"
)

// ExampleHandler 示例处理函数
func ExampleHandler(c *gin.Context) {
    utils.Success(c, map[string]string{"aaa": "8888"},201,"错误0001")
}

func HomeHandler(c * gin.Context) {
	utils.Success(c)
}