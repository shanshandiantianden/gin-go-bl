package v1

import (
	"gin-go-bl/coveralls/models"
	"gin-go-bl/middleware"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data models.User
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}

	user := data.UserName
	password := data.Password
	code := models.CheckLogin(user, password)
	switch {
	case user == "" || password == "":
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码为空",
		})
		return
	case code != 200:

		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码错误!",
		})
		return
	case code == 200:
		data = models.GetUserId(user)
		token, _ := middleware.ReleaseToken(data)
		c.JSON(http.StatusOK, gin.H{
			"userid": data.ID,
			"user":   user,
			"token":  token,
			"code":   utils.SUCCESS,
			"msg":    "登陆成功",
		})

	}

}
