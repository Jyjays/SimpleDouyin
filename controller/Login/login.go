package login

import (
	"net/http"

	"github.com/Jyjays/SimpleDouyin/models"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	var request YourProtoPackage.DouyinUserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 在这里使用 GORM 进行数据库查询，验证登录信息
	var user models.User
	result := db.Where("username = ? AND password = ?", request.Username, request.Password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 登录成功，构造并返回 Protobuf 响应
	response := YourProtoPackage.DouyinUserLoginResponse{
		StatusCode: 0,
		UserId:     user.ID,
		Token:      "your_generated_token",
	}
	c.ProtoBuf(http.StatusOK, &response)
}
