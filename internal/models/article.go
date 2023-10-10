package models

type Article struct {
	BaseModel
	Content    string   `gorm:"longtext;not null" json:"content"`
	Title      string   `gorm:"varchar(20);not null" json:"title"`
	UserUid    uint     `json:"userUid" gorm:"not null"`
	Img        string   `json:"img"`
	CategoryId uint     `gorm:"int;not null" json:"categoryid"`
	Category   Category `gorm:"foreignKey:CategoryId"`
}

func (Article) TableName() string {
	return "article"
}
