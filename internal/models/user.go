package models

import (
	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseModel
	UUID     uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`
	UserName string    `gorm:"varchar(20);not null"  json:"username" validate:"required,min=4,max=12" label:"用户名"`
	NickName string    `gorm:"varchar(40)"  json:"nickName"`
	Phone    string    `gorm:"varchar(20);not null" json:"phone" validate:"required,number=true,min=11,max=11"label:"手机号"`
	Password string    `gorm:"size:255;not null" json:"password" validate:"required,min=4,max=15"label:"密码"`
	Avatar   string    `gorm:"size:255" json:"avatar" `
	Role     int       `gorm:"type:int;colum:role;default:1 ;comment: 权限 0管理员 1用户"`
	R        RoleBase  `gorm:"foreignKey:Role" json:"_"`
}

func (User) TableName() string {
	return "user"
}
