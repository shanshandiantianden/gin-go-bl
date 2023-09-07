package main

import (
	"gin-go-bl/conf"
	"gin-go-bl/internal/database"
	"gin-go-bl/internal/di"
	middlewares2 "gin-go-bl/internal/middlewares"
	"gin-go-bl/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
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
	conf.ConfInit()
	// 初始化数据库连接
	db := database.MysqlServicesInit()
	// 初始化依赖注入容器
	di.InitializeDIContainer(db)
	r.Use(gin.Recovery(), middlewares2.Cors(), middlewares2.Loger(), middlewares2.UnifiedResponseMiddleware())
	router.ApiRouter(r)

	if err := r.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	} // 监听并在 0.0.0.0:8090 上启动服务
}
