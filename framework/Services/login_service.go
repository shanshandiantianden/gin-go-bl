package Services

import (
	"errors"
	"gin-go-bl/framework/Models"
	"gin-go-bl/utils"
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
			return nil, utils.ErrUserNotExist.GetStatusCode()
			//recover()
		}
		log.Println(err)
	}

	if !utils.BcryptCheck(password, user.Password) {
		return nil, utils.ErrUserPassword.GetStatusCode()
	}

	return &user, utils.OK.GetStatusCode()
}
