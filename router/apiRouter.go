package router

import (
	"gin-go-bl/framework/Services"
	"gin-go-bl/framework/api/v1"
	"gin-go-bl/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r1 := r.Group("/v1")
	{

		// 初始化服务层并传递给控制器
		userService := Services.UserServiceImpl{}
		itemController := v1.NewUserController(userService)
		//用户接口

		r1.POST("/user", itemController.RegisterUser)
		r1.GET("/user/:Size/:Page", middlewares.JWTAuth(), itemController.GetAllUser)
		r1.PUT("/user", middlewares.JWTAuth(), itemController.EditMeUserInfo)
		r1.DELETE("/user/:u", itemController.DeleteUser)

		////分类接口
		//r1.POST("/cate", middlewares.AuthMiddleware(), v12.AddCategory)
		//r1.GET("/cate/:num/:size", v12.GetCategory)
		//r1.PUT("/cate/:id", middlewares.AuthMiddleware(), v12.UpdateCategory)
		//r1.DELETE("/cate/:id", middlewares.AuthMiddleware(), v12.DeleteCategory)
		//
		////文章接口
		//r1.POST("/article", middlewares.AuthMiddleware(), v12.AddArticle)
		//r1.GET("/articles/:num/:size", v12.GetAllArticle)
		//r1.GET("/articles", v12.GetAllArticles)
		//r1.GET("/article/:id", v12.GetArticle)
		//r1.GET("/userarticle/:id/:num/:size", v12.GetUserArticle)
		//r1.GET("/myarticle/:num/:size", middlewares.AuthMiddleware(), v12.GetMeArticle)
		//r1.GET("/catearticle/:id/:num/:size", v12.GetCaArticle)
		//r1.PUT("/article/:id", middlewares.AuthMiddleware(), v12.EditArticle)
		//r1.DELETE("/article/:id", middlewares.AuthMiddleware(), v12.DeleteArticle)

		//登录
		r1.POST("/login", itemController.PasswordLogin)

	}

}
