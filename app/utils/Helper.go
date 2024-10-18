package utils

// 统一的JSON格式返回方法
import (
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

// 加密用户密码
func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// 验证密码
func hashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
