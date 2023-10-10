package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey;comment: 主键ID"`
	CreatedAt LocalTime      `gorm:"colum:created_at ;comment: 创建时间"`
	UpdatedAt LocalTime      `gorm:"colum:updated_at ;comment: 更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment: 删除时间"`
}
