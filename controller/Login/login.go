package login

import (
	"net/http"

	"github.com/Jyjays/SimpleDouyin/global"
	"github.com/Jyjays/SimpleDouyin/models"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在这里使用 GORM 进行数据库查询，验证登录信息
	var user models.User
	result := global.Db.Where("username = ? AND password = ?", request.Username, request.Password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 登录成功，构造并返回 JSON 响应
	response := gin.H{
		"status_code": 0,
		"user_id":     user.Id,
		"token":       "your_generated_token",
	}
	c.JSON(http.StatusOK, response)
}
