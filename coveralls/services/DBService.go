package services

import (
	"fmt"
	"gin-go-bl/conf"
	"gin-go-bl/coveralls/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"time"
)

var DB *gorm.DB
var err error

func MysqlServicesInit() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		conf.ConfigObj.User,
		conf.ConfigObj.Password,
		conf.ConfigObj.Host,
		conf.ConfigObj.Port,
		conf.ConfigObj.Database,
		conf.ConfigObj.Charset,
		url.QueryEscape(conf.ConfigObj.Loc))
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	//数据库迁移表，第一次启动后，可以注释掉
	DB.AutoMigrate(&models.User{}, &models.Admin{}, &models.Article{}, &models.Category{})
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}