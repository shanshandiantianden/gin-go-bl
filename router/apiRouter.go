package router

import (
	v12 "gin-go-bl/coveralls/api/v1"
	"gin-go-bl/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine) {
	r1 := r.Group("/v1")
	{
		//用户接口
		r1.POST("/user", v12.AddUser)
		r1.GET("/user/:num/:size", v12.GetUser)
		r1.PUT("/user/:id", middleware.AuthMiddleware(), v12.UpdateUser)
		r1.DELETE("/user/:id", middleware.AuthMiddleware(), v12.DeleteUser)

		//分类接口
		r1.POST("/cate", middleware.AuthMiddleware(), v12.AddCategory)
		r1.GET("/cate/:num/:size", v12.GetCategory)
		r1.PUT("/cate/:id", middleware.AuthMiddleware(), v12.UpdateCategory)
		r1.DELETE("/cate/:id", middleware.AuthMiddleware(), v12.DeleteCategory)

		//文章接口
		r1.POST("/article", middleware.AuthMiddleware(), v12.AddArticle)
		r1.GET("/articles/:num/:size", v12.GetAllArticle)
		r1.GET("/articles", v12.GetAllArticles)
		r1.GET("/article/:id", v12.GetArticle)
		r1.GET("/userarticle/:id/:num/:size", v12.GetUserArticle)
		r1.GET("/myarticle/:num/:size", middleware.AuthMiddleware(), v12.GetMeArticle)
		r1.GET("/catearticle/:id/:num/:size", v12.GetCaArticle)
		r1.PUT("/article/:id", middleware.AuthMiddleware(), v12.EditArticle)
		r1.DELETE("/article/:id", middleware.AuthMiddleware(), v12.DeleteArticle)

		//登录
		r1.POST("/login", v12.Login)

	}

}
