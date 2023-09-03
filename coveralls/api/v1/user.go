package v1

import (
	"gin-go-bl/coveralls/models"
	"gin-go-bl/coveralls/services"
	"gin-go-bl/middlewares"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	var service services.UserService
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	msg, code := middlewares.Validate(&user)

	if code != utils.SUCCESS {
		c.JSON(200, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}
	regUser, code := service.Register(user)
	c.JSON(200, gin.H{
		"data":    regUser,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
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
		"total":   total,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}
func GetMeUser(c *gin.Context) {
	var service services.UserService
	tokenUser, _ := c.Get("user")
	user := tokenUser.(*models.User)
	data, code := service.GetUserInfo(user.UUID)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func UpdateMeUser(c *gin.Context) {

	var service services.UserService
	var updateData models.User
	uuidString, _ := c.Get("user_uuid")
	uuid := uuidString.(uuid.UUID)
	if err := c.ShouldBind(&updateData); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	cCode := service.CheckUser(updateData.UserName)
	if cCode != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  cCode,
			"message": utils.GetErrMsg(cCode),
		})
		return
	}
	if updateData.Password != "" {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "参数错误",
		})
		return

	}

	uCode := service.UpdateUser(uuid, &updateData)
	if uCode != 200 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  cCode,
			"message": utils.GetErrMsg(cCode),
		})
		return
	}
	c.JSON(200, gin.H{
		"data":    updateData,
		"status":  uCode,
		"message": utils.GetErrMsg(uCode),
	})
}

//func UpdateMeUser(c *gin.Context) {
//	tokenUser, _ := c.Get("user")
//	var data models.User
//	//id, _ := strconv.Atoi(c.Param("id"))
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
//func DeleteUser(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Param("id"))
//	data := models.DeleteUser(id)
//	c.JSON(200, gin.H{
//		"status":  data,
//		"message": utils.GetErrMsg(data),
//	})
//}
