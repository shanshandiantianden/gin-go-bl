package services

import (
	"errors"
	"gin-go-bl/coveralls/models"
	"gin-go-bl/utils"
	"gorm.io/gorm"
	"log"
)

type LoginService struct{}

func (loginService LoginService) FindUserInfo(name string, password string) (*models.User, int) {
	var user models.User
	err := DB.Raw("SELECT * FROM user WHERE user_name  = ?", name).First(&user).Error
	//如果err为gorm.ErrRecordNotFound(查询记录为空)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//log.Println(err)
			return nil, utils.ERROR_USERNAME_USED
			//recover()
		}
		log.Println(err)
	}

	if !utils.BcryptCheck(password, user.Password) {
		return nil, utils.ERROR_PASSWORD_WRONG
	}

	return &user, utils.SUCCESS
}
