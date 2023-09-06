package Services

import (
	"gin-go-bl/framework/Models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
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

func (us UserServiceImpl) GetInfo(uuid uuid.UUID) (utils.Error, int) {
	var user Models.User
	err := DB.Raw("SELECT * FROM user WHERE uuid = ?", uuid).Scan(&user).Error
	if err != nil {
		log.Println(err)
		return utils.ErrServer, http.StatusInternalServerError
		//recover()
	}
	return utils.OK.WithData(user), http.StatusOK
}

func (us UserServiceImpl) GetAllInfo(pageSize int, pageNum int) (utils.Error, int) {
	db := DB
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	total := int64(0)
	list := []Models.User{}
	err = db.Model(&Models.User{}).Count(&total).Error
	if err != nil {
		return utils.ErrServer, http.StatusInternalServerError
	}
	err = db.Raw("SELECT * FROM user LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		log.Println(err)
		return utils.ErrServer, http.StatusInternalServerError
	}
	return utils.OK.WithData(map[string]any{
		"list":  list,
		"total": total,
	}), http.StatusOK
}

func (us UserServiceImpl) Create(u Models.User) (utils.Error, int) {
	ok := us.CheckUser(u.UserName)
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	if ok {
		err := DB.Create(&u).Error
		if err != nil {
			return utils.ErrServer, http.StatusInternalServerError
		}
		return utils.OK.WithData(u), http.StatusOK
	}
	return utils.ErrUserExist, http.StatusOK
}

func (us UserServiceImpl) EditInfo(uuid uuid.UUID, data *Models.User) (utils.Error, int) {
	var user Models.User

	//code:= us.CheckUser()

	//data.Password = utils.BcryptHash(data.Password)
	err := DB.Model(&user).Where(
		"uuid = ?",
		uuid,
	).Updates(data).Error
	if err != nil {
		return utils.ErrServer, http.StatusInternalServerError
	}
	return utils.OK.WithData(data), http.StatusOK

}

func (us UserServiceImpl) Delete(uuid uuid.UUID) (utils.Error, int) {
	// 使用 ? 作为参数占位符
	err := DB.Exec("DELETE FROM user WHERE uuid = ?", uuid).Error
	if err != nil {
		// 处理错误并返回适当的错误代码
		log.Println(err)
		return utils.ErrServer, http.StatusInternalServerError
	}
	return utils.OK, http.StatusOK

}
