package handler

import (
	"ggin/app/models"
	"ggin/app/utils"

	"github.com/gin-gonic/gin"
)

// ProfileHandler 获取用户个人信息
func ProfileHandler(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.Error(c, 401, "用户未认证")
		return
	}
	// 将 user 转换为 models.User 类型
	userInfo, ok := user.(map[string]interface{})
	if !ok {
		utils.Error(c, 500, "用户类型转换失败")
		return
	}

	var userDetails models.User
	if err := utils.DB.Select("id,username,account,password").Where("id = ?", userInfo["id"]).First(&userDetails).Error; err != nil {
		utils.Error(c, 500, "获取用户信息失败", err)
		return
	}

	utils.Success(c, gin.H{"user": userDetails})
}