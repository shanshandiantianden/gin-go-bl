package Services

import (
	"errors"
	"gin-go-bl/framework/Models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
)

type UserServiceImpl struct{}

func (us UserServiceImpl) CheckUser(username string) (code int) {
	err := DB.Raw("SELECT id FROM user WHERE user_name  = ?", username).Error
	//如果err不为gorm.ErrRecordNotFound(查询记录为空)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//log.Println(err)
		return utils.ERROR_USERNAME_USED
		//recover()
	}
	return utils.SUCCESS
}

func (us UserServiceImpl) CheckUUID(uuid uuid.UUID) (ok bool) {
	err := DB.Raw("SELECT id FROM user WHERE uuid  = ?", uuid).Error
	//如果err不为gorm.ErrRecordNotFound(查询记录为空)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		//log.Println(err)
		return false
		//recover()
	}
	return true
}

func (us UserServiceImpl) GetUserInfo(uuid uuid.UUID) (user Models.User, code int) {
	err := DB.Raw("SELECT * FROM user WHERE uuid = ?", uuid).Scan(&user).Error
	if err != nil {
		log.Println(err)
		recover()
	}
	return user, utils.SUCCESS
}

func (us UserServiceImpl) GetAllUserInfo(pageSize int, pageNum int) (list []Models.User, total int64, err error) {
	db := DB
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	err = db.Model(&Models.User{}).Count(&total).Error
	if err != nil {
		return
	}
	err = db.Raw("SELECT * FROM user LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		log.Println(err)
		return list, total, err
	}
	return list, total, nil
}

func (us UserServiceImpl) Register(u Models.User) (Models.User, int) {
	code := us.CheckUser(u.UserName)
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

func (us UserServiceImpl) UpdateUser(uuid uuid.UUID, data *Models.User) int {
	var user Models.User

	//code:= us.CheckUser()

	//data.Password = utils.BcryptHash(data.Password)
	err := DB.Model(&user).Where(
		"uuid = ?",
		uuid,
	).Updates(data).Error

	if err != nil {

		return utils.ERROR
	}
	return utils.SUCCESS

}

func (us UserServiceImpl) DeleteUser(uuid uuid.UUID) int {
	// 使用 ? 作为参数占位符
	err := DB.Exec("DELETE FROM user WHERE uuid = ?", uuid).Error
	if err != nil {
		// 处理错误并返回适当的错误代码
		log.Println(err)
		return utils.ERROR
	}
	return utils.SUCCESS

}
