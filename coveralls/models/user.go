package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UUID      uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`
	CreatedAt LocalTime      `json:"created_at"`
	UpdatedAt LocalTime      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"json:"deleted_at"`
	UserName  string         `gorm:"varchar(20);not null"  json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Phone     string         `gorm:"varchar(20);not null" json:"phone" validate:"required,number=true,min=11,max=11"label:"手机号"`
	Password  string         `gorm:"size:255;not null" json:"password" validate:"required,min=4,max=15"label:"密码"`
	Avatar    string         `gorm:"size:255;not null" json:"avatar" `
}

func (User) TableName() string {
	return "user"
}
