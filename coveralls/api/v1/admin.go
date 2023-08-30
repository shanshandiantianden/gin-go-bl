package v1

import (
	"gin-go-bl/coveralls/models"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func Index(c *gin.Context) {
	c.HTML(200, "infologin.html", nil)
}

func AddAdmin(c *gin.Context) {
	var data models.Admin
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	code := models.CreatAdmin(&data)
	c.JSON(200, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
