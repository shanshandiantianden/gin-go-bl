package v1

import (
	"gin-go-bl/framework/Models"
	"gin-go-bl/framework/Services"
	"gin-go-bl/middlewares"
	"gin-go-bl/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginResponse struct {
	Token    string                  `json:"token"` // 用户身份标识
	UserInfo *Models.SessionUserInfo `json:"userinfo"`
}

func (ctrl *UserController) RegisterUser(c *gin.Context) {
	var user Models.User
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
	regUser, code := ctrl.UserService.Register(user)
	c.JSON(200, gin.H{
		"data":    regUser,
		"status":  code,
		"message": utils.GetErrMsg(code),
	})
}

func (ctrl *UserController) PasswordLogin(c *gin.Context) {
	req := new(loginRequest)
	res := new(loginResponse)

	var loginService Services.LoginService
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, "bindJsonFail data is invalid")
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusNotFound, utils.ErrLoginNil)
		return
	}
	info, code := loginService.FindUserInfo(req.Username, req.Password)
	switch {
	case code == utils.ErrUserNotExist.GetStatusCode():
		c.JSON(http.StatusNotFound, utils.ErrUserNotExist)
		return
	case code == utils.ErrUserPassword.GetStatusCode():
		c.JSON(http.StatusNotFound, utils.ErrUserPassword)
		return
	default:
		// 用户信息
		res.UserInfo = &Models.SessionUserInfo{
			UserID:   info.ID,
			UserName: info.UserName,
			UUID:     info.UUID,
		}
		res.Token = utils.CreateToken(c, res.UserInfo)
		c.JSON(http.StatusOK, utils.OK.WithData(res))
	}
}
