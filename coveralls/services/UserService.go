package services

import (
	"fmt"
	"gin-go-bl/coveralls/models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
)

//var NewZzOrmService = &UserService{
//	user: datasource.DefaultDB,
//}

type UserService struct{}

func (userService UserService) CheckUser(username string) (code int) {
	var user models.User
	err := DB.Raw("SELECT id FROM user WHERE user_name  = ?", username).First(&user).Error
	if err != nil {
		panic(err)
	}
	if user.ID != 0 {
		return utils.ERROR_USERNAME_USED
	}
	return utils.SUCCESS
}

func (userService UserService) GetUserInfo(uuid uuid.UUID) (resUser models.User, code int) {
	var user models.User
	err := DB.Raw("SELECT * FROM user WHERE uuid = ?", uuid).Scan(&user).Error
	if err != nil {
		panic(err)
	}
	return user, utils.SUCCESS
}

func (userService UserService) GetAllUserInfo(pageSize int, pageNum int) (list interface{}, total int64, err error) {
	var userList []models.User
	db := DB
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	err = db.Model(&models.User{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Raw("SELECT * FROM user LIMIT ?,?", offset, limit).Scan(&userList).Error
	if err != nil {
		fmt.Println(err)
		return userList, total, err
	}
	//err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&user).Error
	return userList, total, nil
}

func (userService UserService) Register(u models.User) (models.User, int) {
	userService.CheckUser(u.UserName)
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err := DB.Create(&u).Error
	if err != nil {
		return u, utils.ERROR
	}
	return u, utils.SUCCESS

}

//
//func DeleteUser(id int) int {
//	var user User
//	err := DB.Where("id = ?", id).Delete(&user).Error
//
//	if err != nil {
//
//		return utils.ERROR
//	}
//	return utils.SUCCESS
//
//}
//
//func UpdateUser(id int, data *User) int {
//	var user User
//	var up = make(map[string]interface{})
//	up["user_name"] = data.UserName
//	up["phone"] = data.Phone
//	up["avatar"] = data.Avatar
//
//	data.Password = utils.ScryptPassword(data.Password)
//	err := DB.Model(&user).Where("id = ?", id).Updates(data).Error
//
//	if err != nil {
//
//		return utils.ERROR
//	}
//	return utils.SUCCESS
//
//}
//
//func GetUserId(name string) User {
//	var user User
//	DB.Where("user_name =?", name).First(&user)
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return user
//	}
//	return user
//}
//
//func CheckLogin(name string, password string) (code int) {
//	var user User
//	DB.Where("user_name = ?", name).First(&user)
//	if user.ID == 0 {
//		code = utils.ERROR_USERNAME_NOT
//		return
//	}
//	if utils.ScryptPassword(password) != user.Password {
//		code = utils.ERROR_PASSWORD_WRONG
//		return
//	}
//	code = utils.SUCCESS
//	return
//}
