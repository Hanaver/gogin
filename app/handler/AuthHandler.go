package handler

import (
	"ggin/app/models"

	"ggin/app/utils"

	"github.com/gin-gonic/gin"

	"fmt"
)

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

// LoginHandler 处理用户登录请求
func LoginHandler(c *gin.Context) {
	var loginData struct {
		Account  string `json:"account" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// 绑定请求体到登录数据结构
	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.Error(c, 400, "无效的请求数据", err)
		return
	}

	// 查找用户
	var user models.User
	if err := utils.DB.Select("id, account, password").Where("account = ?", loginData.Account).First(&user).Error; err != nil {
		utils.Error(c, 401, "账号不存在", err)
		return
	}

	// 验证密码
	if !utils.HashPassword(loginData.Password, user.Password) {
		utils.Error(c, 401, "账号或密码错误")
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.Error(c, 500, "令牌生成失败", err)
		return
	}

	utils.Success(c, gin.H{"token": token})
}


// userExists 检查用户名或账号是否已存在
func userExists(account string) bool {
	var user models.User
	utils.DB.Where("account = ?", account).First(&user)
	return user.ID != 0
}