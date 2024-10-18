package utils

import (
	"ggin/app/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 是全局的数据库连接
var DB *gorm.DB


// InitDB 初始化数据库连接
func InitDB() {
	var err error
	// 数据库连接信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/gogin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	// 自动迁移
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
