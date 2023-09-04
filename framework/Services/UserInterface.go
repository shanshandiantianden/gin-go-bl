package Services

import (
	"gin-go-bl/framework/Models"
	uuid "github.com/satori/go.uuid"
)

type UserInterface interface {
	// CheckUser 检查用户是否存在
	CheckUser(username string) (code int)

	// GetUserInfo 根据UUID获取用户信息
	GetUserInfo(uuid uuid.UUID) (user Models.User, code int)

	// GetAllUserInfo 分页获取所有用户
	GetAllUserInfo(pageSize int, pageNum int) (list interface{}, total int64, err error)
	// Register 用户注册
	Register(u Models.User) (Models.User, int)

	// UpdateUser 更新用户信息
	UpdateUser(uuid uuid.UUID, data *Models.User) int

	DeleteUser(uuid uuid.UUID) int
}
