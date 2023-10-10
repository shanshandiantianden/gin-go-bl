package database

import (
	"fmt"
	"gin-go-bl/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"os"
	"time"
)

func MysqlServicesInit() *gorm.DB {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		conf.ConfigObj.User,
		conf.ConfigObj.Password,
		conf.ConfigObj.Host,
		conf.ConfigObj.Port,
		conf.ConfigObj.Database,
		conf.ConfigObj.Charset,
		url.QueryEscape(conf.ConfigObj.Loc))
	// 连接数据库

	NewLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢查询阈值
		LogLevel:      logger.Info, //log lever
		Colorful:      true,        //禁用彩色打印
	})

	DB, err := gorm.Open(mysql.Open(args), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: NewLogger,
	})
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	////数据库迁移表，第一次启动后，可以注释掉
	//err = DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Category{})
	//if err != nil {
	//	panic("failed to migrate: " + err.Error())
	//}
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return DB
}
