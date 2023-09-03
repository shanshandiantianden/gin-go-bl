package services

import (
	"errors"
	"fmt"
	"gin-go-bl/coveralls/models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
)

//var NewZzOrmService = &UserService{
//	user: datasource.DefaultDB,
//}

type UserService struct{}

func (userService UserService) CheckUser(username string) (code int) {
	var user models.User
	err := DB.Raw("SELECT id FROM user WHERE user_name  = ?", username).First(&user).Error

	//如果err不为gorm.ErrRecordNotFound(查询记录为空)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//log.Println(err)
		return utils.ERROR_USERNAME_USED
		//recover()
	}
	return utils.SUCCESS
}

func (userService UserService) GetUserInfo(uuid uuid.UUID) (user models.User, code int) {
	err := DB.Raw("SELECT * FROM user WHERE uuid = ?", uuid).Scan(&user).Error
	if err != nil {
		log.Println(err)
		recover()
	}
	return user, utils.SUCCESS
}

func (userService UserService) GetAllUserInfo(pageSize int, pageNum int) (list interface{}, total int64, err error) {
	db := DB
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	err = db.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Raw("SELECT * FROM user LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		fmt.Println(err)
		return list, total, err
	}
	return list, total, nil
}

func (userService UserService) Register(u models.User) (models.User, int) {
	code := userService.CheckUser(u.UserName)
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	if code == 200 {
		err := DB.Create(&u).Error
		if err != nil {
			return u, utils.ERROR
		}
		return u, utils.SUCCESS
	}
	return u, code
}

func (userService UserService) UpdateUser(uuid uuid.UUID, data *models.User) int {
	var user models.User

	//data.Password = utils.BcryptHash(data.Password)
	err := DB.Model(&user).Where("uuid = ?", uuid).Updates(data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func (userService UserService) DeleteUser(uuid uuid.UUID) int {
	var user models.User
	err := DB.Where("uuid = ?", uuid).Delete(&user).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}
