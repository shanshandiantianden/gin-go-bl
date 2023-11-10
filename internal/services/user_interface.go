package services

import (
	"gin-go-bl/internal/errmsg"
	"gin-go-bl/internal/models"
	uuid "github.com/satori/go.uuid"
)

type CRUDOperation func(interface{}) (errmsg.Error, int)

type UserServiceInterface interface {
	CRUDService
	// CheckUser 检查用户是否存在
	CheckUserName(username string) (ok bool)
	// CheckUUID
	CheckUUID(uuid uuid.UUID) (ok bool)
	FindUserInfo(name string, password string) (*models.User, int)
}
