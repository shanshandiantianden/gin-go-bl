package main

import (
	"gin-go-bl/conf"
	"gin-go-bl/internal/database"
	"gin-go-bl/internal/di"
	"gin-go-bl/internal/router"
	"log"
)

func main() {

	conf.ConfInit()
	// 初始化数据库连接
	db := database.MysqlServicesInit()
	// 初始化依赖注入容器
	di.InitializeDIContainer(db)

	r := router.ApiRouter()
	if err := r.Run(":8090"); err != nil {
		log.Fatal(err.Error())
	} // 监听并在 0.0.0.0:8090 上启动服务
}
