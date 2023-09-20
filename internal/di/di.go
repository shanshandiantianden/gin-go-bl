package di

import (
	"gin-go-bl/internal/services"
	"gorm.io/gorm"
)

var Container *DIContainer

type DIContainer struct {
	UserService    services.UserServiceInterface
	ArticleService services.ArticleServiceInterface
}

func NewDIContainer(db *gorm.DB) *DIContainer {
	return &DIContainer{
		UserService:    services.NewUserService(db),
		ArticleService: services.NewArticleService(db),
	}
}

func InitializeDIContainer(db *gorm.DB) {
	Container = NewDIContainer(db)
}
