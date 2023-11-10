package router

import (
	"gin-go-bl/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	//debugger模式
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//r.LoadHTMLGlob("templates/*")
	////r.LoadHTMLGlob("views/*")
	////r.LoadHTMLGlob("views/**/*")
	//r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"index": "默认",
		})
	})

	r.Use(gin.Recovery(), middlewares.Cors(), middlewares.Loger())

	// 设置路由和中间件
	// 这里你可以按照建议进行路由组织和中间件的抽象
	UserRouter(r)

	return r
}
