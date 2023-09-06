package Services

import (
	"errors"
	"gin-go-bl/internal/Models"
	"gin-go-bl/internal/errmsg"
	utils2 "gin-go-bl/pkg/utils"
	"gorm.io/gorm"
	"log"
)

type LoginService struct{}

func (loginService LoginService) FindUserInfo(name string, password string) (*Models.User, int) {
	var user Models.User
	err := DB.Raw("SELECT * FROM user WHERE user_name  = ?", name).First(&user).Error
	//如果err为gorm.ErrRecordNotFound(查询记录为空)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//log.Println(err)
			return nil, errmsg.ErrUserNotExist.GetStatusCode()
			//recover()
		}
		log.Println(err)
	}

	if !utils2.BcryptCheck(password, user.Password) {
		return nil, errmsg.ErrUserPassword.GetStatusCode()
	}

	return &user, errmsg.OK.GetStatusCode()
}
