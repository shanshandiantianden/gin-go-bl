package v1

import (
	"fmt"
	"gin-go-bl/global"
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/models"
	"gin-go-bl/internal/services"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{
		UserService: userService,
	}
}
func (ctrl *UserController) GetAllUser(c *gin.Context) {
	size, _ := strconv.Atoi(c.Param("Size"))
	page, _ := strconv.Atoi(c.Param("Page"))

	if size == 0 {
		size = global.DefaultPageSize
	}
	if page == 0 {
		page = global.DefaultPage
	}
	res, httpStatus := ctrl.UserService.GetAllInfo(size, page)
	c.JSON(httpStatus, res)
}
func (ctrl *UserController) GetMeUser(c *gin.Context) {
	tokenUser, exists := c.Get("user")
	if !exists {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, errmsg.ErrUnauthorized.WithData(nil))
		return
	}
	user, ok := tokenUser.(*models.User)
	if !ok {
		// 返回未经授权的错误响应
		c.JSON(http.StatusUnauthorized, errmsg.ErrUnauthorized.WithData(nil))
		return
	}
	res, httpStatus := ctrl.UserService.GetInfo(user.UUID)
	c.JSON(httpStatus, res)
}

func (ctrl *UserController) EditMeUserInfo(c *gin.Context) {
	var updateData models.User
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
	cOk := ctrl.UserService.CheckUserName(updateData.UserName)

	if !cOk {
		// 返回用户名冲突的错误响应
		c.JSON(http.StatusUnauthorized, errmsg.ErrUserExist.WithData(nil))
		return
	}
	// 检查密码是否为空
	if updateData.Password != "" {

		c.JSON(http.StatusUnauthorized, errmsg.ErrParam.WithData(nil))
		return

	}

	res, httpStatus := ctrl.UserService.EditInfo(uuid, &updateData)
	c.JSON(httpStatus, res)
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
		c.JSON(http.StatusUnauthorized, errmsg.ErrUserNotExist.WithData(nil))
		return
	}
	res, httpStatus := ctrl.UserService.Delete(u)
	fmt.Println(res.ToString())
	c.JSON(httpStatus, res)
}
