package Services

import (
	uuid "github.com/satori/go.uuid"
)

type UserService interface {
	CRUDService
	// CheckUser 检查用户是否存在
	CheckUser(username string) (ok bool)
	// CheckUUID
	CheckUUID(uuid uuid.UUID) (ok bool)
}
