package v1

import (
	"gin-go-bl/coveralls/services"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// import (
//
//	"gin-go-bl/coveralls/models"
//	"gin-go-bl/coveralls/services"
//	"gin-go-bl/middleware"
//	"gin-go-bl/utils"
//	"github.com/gin-gonic/gin"
//	"strconv"
//
// )
//
//	func AddUser(c *gin.Context) {
//		var data models.User
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(200, "bindJsonFail data is invalid")
//			return
//		}
//		msg, code := middleware.Validate(&data)
//
//		if code != utils.SUCCESS {
//			c.JSON(200, gin.H{
//				"status":  code,
//				"message": msg,
//			})
//			return
//		}
//
//		code = data.CheckUser(data.UserName)
//		if code == 200 {
//			data.CreatUser()
//			//models.CreatUser(&data)
//		}
//		c.JSON(200, gin.H{
//			"data":    data,
//			"status":  code,
//			"message": utils.GetErrMsg(code),
//		})
//	}
func GetAllUser(c *gin.Context) {
	var service services.UserService
	size, _ := strconv.Atoi(c.Param("Size"))
	page, _ := strconv.Atoi(c.Param("Page"))
	if size == 0 {
		size = -1
	}
	if page == 0 {
		page = -1
	}

	data, total, err := service.GetAllUserInfo(size, page)
	if err != nil {
		log.Println(err)
	}
	code := utils.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"size":    total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

//
//func UpdateUser(c *gin.Context) {
//	user, _ := c.Get("user")
//	var data models.User
//	id, _ := strconv.Atoi(c.Param("id"))
//	if err := c.ShouldBind(&data); err != nil {
//		c.JSON(200, "bindJsonFail data is invalid")
//		return
//	}
//	code := data.CheckUser(data.UserName)
//
//	switch {
//	case code == 200:
//		models.UpdateUser(id, &data)
//	case code == 1001:
//
//		if uint(id) == user.(models.User).ID {
//			models.UpdateUser(id, &data)
//		}
//
//	case 1 == user.(models.User).ID:
//
//	}
//	c.JSON(200, gin.H{
//		"data":    data,
//		"status":  code,
//		"message": utils.GetErrMsg(code),
//	})
//}
//
//func DeleteUser(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	data := models.DeleteUser(id)
//	c.JSON(200, gin.H{
//		"status":  data,
//		"message": utils.GetErrMsg(data),
//	})
//}
