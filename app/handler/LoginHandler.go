package handler

import (
	"ggin/app/models"

	"ggin/app/utils"

	"github.com/gin-gonic/gin"

	"fmt"
)

func LoginHandler(c *gin.Context) {
	
}

// RegisterHandler 处理用户注册请求
func RegisterHandler(c *gin.Context) {
	var user models.User
	
	// 绑定请求体到用户结构
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Error(c, "无效的请求数据", 400)
		return
	}
	
	// 检查用户名和账号是否已存在
	if userExists(user.Account) {
		print("1111111111")
		
		utils.Error(c, "账号已存在", 409)
		return
	}
	
	// 对密码进行哈希处理
	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		utils.Error(c, "密码处理失败", 500)
		return
	}
	user.Password = hashedPassword
	
	// 创建新用户
	if err := utils.DB.Create(&user).Error; err != nil {
		utils.Error(c, "用户创建失败", 500)
		return
	}

	fmt.Println("Hello, World!")
	
	utils.Success(c, gin.H{"message": "用户注册成功"})
}

// userExists 检查用户名或账号是否已存在
func userExists(account string) bool {
	var user models.User
	utils.DB.Where("account = ?", account).First(&user)
	return user.ID != 0
}

