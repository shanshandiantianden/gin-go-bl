package models

type RoleBase struct {
	BaseModel
	Title string `json:"title" gorm:"size:20;not null"`
	Pwd   string `json:"pwd" gorm:"size:255"`
	Is    bool   `json:"is" gorm:"column:is"`
}
