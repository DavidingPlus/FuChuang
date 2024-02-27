package users

import (
	"backend/jwt"
	"backend/model/user"
	"backend/response"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*user.UserInfos, 10)
	data = user.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

func DelteUser(c *gin.Context) {
	id := c.Query("id")
	user.DelteUser(id)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// RefreshToken 刷新 Access Token
func RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"code":  1,
			"token": token,
		})
	}
}
