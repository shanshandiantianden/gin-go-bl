package services

import (
	"fmt"
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/models"
)

func (us *UserServiceImpl) FindUserInfo(name string, password string) (*models.User, int) {
	var user models.User
	//err := us.db.Raw("SELECT * FROM user WHERE user_name  = ?", name).Scan(&user).Error
	us.db.Where("user_name = ?", name).First(&user)
	fmt.Println(user)
	////如果err为gorm.ErrRecordNotFound(查询记录为空)
	//fmt.Println(err)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		//log.Println(err)
	//		return nil, errmsg.ErrUserNotExist.GetStatusCode()
	//		//recover()
	//	}
	//	log.Println(err)
	//}
	//
	//if !utils2.BcryptCheck(password, user.Password) {
	//	return nil, errmsg.ErrUserPassword.GetStatusCode()
	//}
	//
	return &user, errmsg.OK.GetStatusCode()
}
