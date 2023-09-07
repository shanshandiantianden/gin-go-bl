package Services

import (
	"gin-go-bl/internal/Models"
	"gin-go-bl/internal/errmsg"
	utils2 "gin-go-bl/pkg/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type UserServiceImpl struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

//func (us *UserServiceImpl) InitializeDatabase(db *gorm.DB) {
//	us.db = db
//}

func (us *UserServiceImpl) CheckUser(username string) (ok bool) {

	var count int

	// 执行原生 SQL 查询
	err := us.db.Raw("SELECT COUNT(*) FROM user WHERE user_name = ?", username).Scan(&count).Error
	if err != nil {
		log.Println(err)
		return false // 返回错误信息
	}

	// 如果记录数量大于 0，则表示 UUID 存在
	return count > 0
}

func (us *UserServiceImpl) CheckUUID(uuid uuid.UUID) (ok bool) {
	var count int

	// 执行原生 SQL 查询
	err := us.db.Raw("SELECT COUNT(*) FROM user WHERE uuid = ?", uuid).Scan(&count).Error
	if err != nil {
		log.Println(err)
		return false // 返回错误信息
	}

	// 如果记录数量大于 0，则表示 UUID 存在
	return count > 0
}

func (us *UserServiceImpl) GetInfo(uuid any) (errmsg.Error, int) {
	var user Models.User
	err := us.db.Raw("SELECT * FROM user WHERE uuid = ?", uuid).Scan(&user).Error
	if err != nil {
		log.Println(err)
		return errmsg.ErrServer, http.StatusInternalServerError
		//recover()
	}
	return errmsg.OK.WithData(user), http.StatusOK
}

func (us *UserServiceImpl) GetAllInfo(pageSize int, pageNum int) (errmsg.Error, int) {
	db := us.db
	limit := pageSize
	offset := pageSize * (pageNum - 1)
	total := int64(0)
	list := []Models.User{}
	err := db.Model(&Models.User{}).Count(&total).Error
	if err != nil {
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	err = db.Raw("SELECT * FROM user LIMIT ?,?", offset, limit).Scan(&list).Error
	if err != nil {
		log.Println(err)
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK.WithData(map[string]any{
		"list":  list,
		"total": total,
	}), http.StatusOK
}

func (us *UserServiceImpl) Create(user any) (errmsg.Error, int) {
	u := user.(Models.User)
	ok := us.CheckUser(u.UserName)
	u.Password = utils2.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	if ok {
		err := us.db.Create(&u).Error
		if err != nil {
			return errmsg.ErrServer, http.StatusInternalServerError
		}
		return errmsg.OK.WithData(u), http.StatusOK
	}
	return errmsg.ErrUserExist, http.StatusOK
}

func (us *UserServiceImpl) EditInfo(uid any, data any) (errmsg.Error, int) {
	var user Models.User
	uid = uid.(uuid.UUID)
	//code:= us.CheckUser()

	//data.Password = utils.BcryptHash(data.Password)
	err := us.db.Model(&user).Where(
		"uuid = ?",
		uid,
	).Updates(data).Error
	if err != nil {
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK.WithData(data), http.StatusOK

}

func (us *UserServiceImpl) Delete(uuid any) (errmsg.Error, int) {
	// 使用 ? 作为参数占位符
	err := us.db.Exec("DELETE FROM user WHERE uuid = ?", uuid).Error
	if err != nil {
		// 处理错误并返回适当的错误代码
		log.Println(err)
		return errmsg.ErrServer, http.StatusInternalServerError
	}
	return errmsg.OK, http.StatusOK

}
