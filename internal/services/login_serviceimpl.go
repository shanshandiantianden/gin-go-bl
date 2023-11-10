package services

import (
	"errors"
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/models"
	utils2 "gin-go-bl/pkg/utils"
	"gorm.io/gorm"
	"log"
)

func (us *UserServiceImpl) FindUserInfo(name string, password string) (u *models.User, code int) {
	var user models.User
	err := us.db.Raw("select * from user where user_name = ?", name).First(&user).Error
	////如果err为gorm.ErrRecordNotFound(查询记录为空)
	//fmt.Println(err)
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
