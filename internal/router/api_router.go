package router

import (
	v1 "gin-go-bl/api/v1"
	"gin-go-bl/internal/di"
	"gin-go-bl/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	r1 := r.Group("/v1")
	{
		//用户接口
		userService := di.Container.UserService
		itemController := v1.NewUserController(userService)
		r1.POST("/user", itemController.RegisterUser)
		r1.GET("/user/:Size/:Page", middlewares.JWTAuth(), itemController.GetAllUser)
		r1.PUT("/user", middlewares.JWTAuth(), itemController.EditMeUserInfo)
		r1.DELETE("/user/:u", itemController.DeleteUser)

		//登录
		r1.POST("/login", itemController.PasswordLogin)

	}

}

func ArticleRouter(r *gin.Engine) {
	r1 := r.Group("/v1")
	{
		////分类接口
		r1.POST("/cate", middlewares.JWTAuth())
		//r2.GET("/cate/:num/:size", v12.GetCategory)
		//r2.PUT("/cate/:id", middlewares.AuthMiddleware(), v12.UpdateCategory)
		//r2.DELETE("/cate/:id", middlewares.AuthMiddleware(), v12.DeleteCategory)
		////
		//////文章接口
		//r2.POST("/article", middlewares.AuthMiddleware(), v12.AddArticle)
		//r2.GET("/articles/:num/:size", v12.GetAllArticle)
		//r2.GET("/articles", v12.GetAllArticles)
		//r2.GET("/article/:id", v12.GetArticle)
		//r2.GET("/userarticle/:id/:num/:size", v12.GetUserArticle)
		//r2.GET("/myarticle/:num/:size", middlewares.AuthMiddleware(), v12.GetMeArticle)
		//r2.GET("/catearticle/:id/:num/:size", v12.GetCaArticle)
		//r2.PUT("/article/:id", middlewares.AuthMiddleware(), v12.EditArticle)
		//r2.DELETE("/article/:id", middlewares.AuthMiddleware(), v12.DeleteArticle)

	}

}
