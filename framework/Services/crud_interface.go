package Services

import (
	"gin-go-bl/framework/Models"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
)

type CRUDService interface {
	GetInfo(uuid uuid.UUID) (utils.Error, int)
	GetAllInfo(pageSize int, pageNum int) (utils.Error, int)
	Create(data Models.User) (utils.Error, int)
	EditInfo(uuid uuid.UUID, data *Models.User) (utils.Error, int)
	Delete(uuid uuid.UUID) (utils.Error, int)
}
