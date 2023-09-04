package Services

import (
	"gin-go-bl/framework/Models"
	uuid "github.com/satori/go.uuid"
)

type UserService interface {
	// CheckUser 检查用户是否存在
	CheckUser(username string) (code int)
	// CheckUUID
	CheckUUID(uuid uuid.UUID) (ok bool)

	// GetUserInfo 根据UUID获取用户信息
	GetUserInfo(uuid uuid.UUID) (user Models.User, code int)

	// GetAllUserInfo 分页获取所有用户
	GetAllUserInfo(pageSize int, pageNum int) (list []Models.User, total int64, err error)
	// Register 用户注册
	Register(u Models.User) (Models.User, int)

	// UpdateUser 更新用户信息
	UpdateUser(uuid uuid.UUID, data *Models.User) int
	//删除用户
	DeleteUser(uuid uuid.UUID) int
}
