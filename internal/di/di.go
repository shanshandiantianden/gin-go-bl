package di

import (
	"gin-go-bl/internal/Services"
	"gorm.io/gorm"
)

var Container *DIContainer

type DIContainer struct {
	UserService    Services.UserServiceInterface
	ArticleService Services.ArticleServiceInterface
}

func NewDIContainer(db *gorm.DB) *DIContainer {
	return &DIContainer{
		UserService:    Services.NewUserService(db),
		ArticleService: Services.NewArticleService(db),
	}
}

func InitializeDIContainer(db *gorm.DB) {
	Container = NewDIContainer(db)
}
