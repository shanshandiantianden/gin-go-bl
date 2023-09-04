package v1

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func Index(c *gin.Context) {
	c.HTML(200, "infologin.html", nil)
}

//func AddAdmin(c *gin.Context) {
//	var data Models.Admin
//	if err := c.ShouldBind(&data); err != nil {
//		c.JSON(200, "bindJsonFail data is invalid")
//		return
//	}
//	code := Models.CreatAdmin(&data)
//	c.JSON(200, gin.H{
//		"data":    data,
//		"status":  code,
//		"message": utils.GetErrMsg(code),
//	})
//}
