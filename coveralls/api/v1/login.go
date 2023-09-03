package v1

import (
	"gin-go-bl/coveralls/models"
	"gin-go-bl/coveralls/services"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PasswordLogin(c *gin.Context) {
	var formdata models.PasswordLogin
	var loginService services.LoginService
	if err := c.ShouldBind(&formdata); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	if formdata.UserName == "" || formdata.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码为空",
		})
	}
	user, code := loginService.FindUserInfo(formdata.UserName, formdata.Password)
	if code == 1001 || code == 1002 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "用户名或密码错误!",
		})
		return
	}

	token := utils.CreateToken(c, user.ID, user.UserName, user.UUID)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"name":   user.UserName,
			"uuid":   user.UUID,
			"avatar": user.Avatar,
		},
		"token": token,
		"code":  utils.SUCCESS,
		"msg":   "登陆成功",
	})

}
