package models

type Article struct {
	Category   Category   `gorm:"foreignKey:CategoryId"`
	Aid        uint       `gorm:"int;not null;primary_key" json:"aid"`
	Title      string     `gorm:"varchar(20);not null" json:"title"`
	CategoryId uint       `gorm:"int;not null" json:"categoryid"`
	UserId     uint       `json:"user_id" gorm:"not null"`
	Desc       string     `gorm:"varchar(255);not null" json:"desc"`
	Content    string     `gorm:"longtext;not null" json:"content"`
	Img        string     `json:"img"`
	CreatedAt  *LocalTime `json:"crtime"`
}

func (Article) TableName() string {
	return "article"
}
