package v1

import (
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/middlewares"
	"gin-go-bl/internal/models"
	"gin-go-bl/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginResponse struct {
	Token    string                  `json:"token"` // 用户身份标识
	UserInfo *models.SessionUserInfo `json:"userinfo"`
}

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	msg, code := middlewares.Validate(&user)
	if code != http.StatusOK {
		c.JSON(http.StatusUnauthorized, errmsg.ErrParam.WithData(msg))
		return
	}
	res, httpStatus := ctrl.UserService.Create(user)
	c.JSON(httpStatus, res)
}

func (ctrl *UserController) PasswordLogin(c *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)

	var loginService services.UserServiceImpl
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusNotFound, errmsg.ErrLoginNil)
		return
	}
	info, code := loginService.FindUserInfo(req.Username, req.Password)
	switch {
	case code == errmsg.ErrUserNotExist.GetStatusCode():
		c.JSON(http.StatusNotFound, errmsg.ErrUserNotExist)
		return
	case code == errmsg.ErrUserPassword.GetStatusCode():
		c.JSON(http.StatusNotFound, errmsg.ErrUserPassword)
		return
	default:
		// 用户信息
		res.UserInfo = &models.SessionUserInfo{
			UserID:   info.ID,
			UserName: info.UserName,
			UUID:     info.UUID,
		}
		res.Token = middlewares.CreateToken(c, res.UserInfo)
		c.JSON(http.StatusOK, errmsg.OK.WithData(res))
	}
}
