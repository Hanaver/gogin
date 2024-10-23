package utils

// 统一的JSON格式返回方法
import (
	"time"

	"ggin/app/models" // 确保替换为实际的项目路径

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// JSONResponse 统一的JSON格式返回方法
func Success(c *gin.Context, data ...interface{}) {
	code := 200
	message := "Successful"
	var responseData interface{} = struct{}{}

	if len(data) > 0 {
		responseData = data[0]
	}
	if len(data) > 1 {
		if codeVal, ok := data[1].(int); ok {
			code = codeVal
		}
	}
	if len(data) > 2 {
		if msgVal, ok := data[2].(string); ok {
			message = msgVal
		}
	}

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    responseData,
	})
}

func Error(c *gin.Context, data ...interface{}) {
	code := 400
	message := "Fail"
	var responseData interface{} = struct{}{}

	if len(data) > 0 {
		if codeVal, ok := data[0].(int); ok {
			code = codeVal
		}
	}
	if len(data) > 1 {
		if msgVal, ok := data[1].(string); ok {
			message = msgVal
		}
	}
	if len(data) > 2 {
		responseData = data[2]
	}

	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"error":    responseData,
	})
}

// 加密用户密码
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func HashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// 生成用户登录令牌
func GenerateJWT(user models.User) (string, error) {
	var userTokenInfo struct {
		Account  string `json:"account" binding:"required"`
		ID uint `json:"id" binding:"required"`
	}
	userTokenInfo.Account = user.Account
	userTokenInfo.ID = user.ID
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userTokenInfo,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	
	tokenString, err := token.SignedString([]byte("AAAAAA"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
