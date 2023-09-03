package services

import (
	"gin-go-bl/coveralls/models"
	uuid "github.com/satori/go.uuid"
)

type UserInterface interface {
	// CheckUser 检查用户是否存在
	CheckUser(username string) (code int)

	// GetUserInfo 根据UUID获取用户信息
	GetUserInfo(uuid uuid.UUID) (user models.User, code int)

	// GetAllUserInfo 分页获取所有用户
	GetAllUserInfo(pageSize int, pageNum int) (list interface{}, total int64, err error)
	// Register 用户注册
	Register(u models.User) (models.User, int)

	// UpdateUser 更新用户信息
	UpdateUser(uuid uuid.UUID, data *models.User) int

	DeleteUser(uuid uuid.UUID) int
}
