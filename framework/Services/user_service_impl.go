package Services

import (
	"gin-go-bl/framework/Models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
	"log"
)

type UserServiceImpl struct{}

func (us UserServiceImpl) CheckUser(username string) (ok bool) {

	var count int

	// 执行原生 SQL 查询
	err := DB.Raw("SELECT COUNT(*) FROM user WHERE user_name = ?", username).Scan(&count).Error
	if err != nil {
		log.Println(err)
		return false // 返回错误信息
	}

	// 如果记录数量大于 0，则表示 UUID 存在
	return count > 0
}

func (us UserServiceImpl) CheckUUID(uuid uuid.UUID) (ok bool) {
	var count int

	// 执行原生 SQL 查询
	err := DB.Raw("SELECT COUNT(*) FROM user WHERE uuid = ?", uuid).Scan(&count).Error
	if err != nil {
		log.Println(err)
		return false // 返回错误信息
	}

	// 如果记录数量大于 0，则表示 UUID 存在
	return count > 0
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

func (us UserServiceImpl) Register(u Models.User) utils.Error {
	ok := us.CheckUser(u.UserName)
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	if ok {
		err := DB.Create(&u).Error
		if err != nil {
			return utils.ErrGormQuery
		}
		return utils.OK.WithData(u)
	}
	return utils.ErrServer
}

func (us UserServiceImpl) UpdateUser(uuid uuid.UUID, data *Models.User) utils.Error {
	var user Models.User

	//code:= us.CheckUser()

	//data.Password = utils.BcryptHash(data.Password)
	err := DB.Model(&user).Where(
		"uuid = ?",
		uuid,
	).Updates(data).Error
	if err != nil {
		return utils.ErrGormQuery
	}
	return utils.OK.WithData(data)

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
