package v1

import (
	"gin-go-bl/framework/Models"
	"gin-go-bl/framework/Services"
	"gin-go-bl/global"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService Services.UserService
}

func NewUserController(userService Services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
func (ctrl *UserController) GetAllUser(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("Size"))
	page, _ := strconv.Atoi(c.Param("Page"))

	if size == 0 {
		size = global.DefaultSize
	}
	if page == 0 {
		page = global.DefaultPage
	}
	data, total, err := ctrl.UserService.GetAllUserInfo(size, page)
	if err != nil {
		log.Println(err)
		// 返回适当的错误响应
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"total":   total,
		"status":  utils.SUCCESS,
		"message": utils.GetErrMsg(utils.SUCCESS),
	})
}
func (ctrl *UserController) GetMeUser(c *gin.Context) {
	tokenUser, exists := c.Get("user")
	if !exists {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	user, ok := tokenUser.(*Models.User)
	if !ok {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	data, code := ctrl.UserService.GetUserInfo(user.UUID)
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func (ctrl *UserController) UpdateMeUser(c *gin.Context) {
	var updateData Models.User
	uuidString, exists := c.Get("user_uuid")
	if !exists {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	uuid, ok := uuidString.(uuid.UUID)
	if !ok {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	if err := c.ShouldBind(&updateData); err != nil {
		// 返回数据无效的错误响应
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	cCode := ctrl.UserService.CheckUser(updateData.UserName)

	if cCode != 200 {
		// 返回用户名冲突的错误响应
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  cCode,
			"message": utils.GetErrMsg(cCode),
		})
		return
	}
	// 检查密码是否为空
	if updateData.Password != "" {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "参数错误",
		})
		return

	}

	uCode := ctrl.UserService.UpdateUser(uuid, &updateData)
	if uCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{
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

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("u"))
	uuidString := c.Param("u")
	var u uuid.UUID
	err := u.UnmarshalText([]byte(uuidString))
	if err != nil {
		return
	}
	if !ctrl.UserService.CheckUUID(u) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "当前用户不存在",
		})
		return
	}
	data := ctrl.UserService.DeleteUser(u)
	c.JSON(200, gin.H{
		"status":  data,
		"message": utils.GetErrMsg(data),
	})
}
