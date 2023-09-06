package Services

import (
	"gin-go-bl/internal/Models"
	"gin-go-bl/internal/errmsg"
	uuid "github.com/satori/go.uuid"
)

type CRUDService interface {
	GetInfo(uuid uuid.UUID) (errmsg.Error, int)
	GetAllInfo(pageSize int, pageNum int) (errmsg.Error, int)
	Create(data Models.User) (errmsg.Error, int)
	EditInfo(uuid uuid.UUID, data *Models.User) (errmsg.Error, int)
	Delete(uuid uuid.UUID) (errmsg.Error, int)
}
