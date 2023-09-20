package services

import (
	"gin-go-bl/internal/errmsg"
)

type CRUDService interface {
	GetInfo(interface{}) (errmsg.Error, int)
	GetAllInfo(pageSize int, pageNum int) (errmsg.Error, int)
	Create(interface{}) (errmsg.Error, int)
	EditInfo(interface{}, interface{}) (errmsg.Error, int)
	Delete(interface{}) (errmsg.Error, int)
}
