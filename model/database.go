package model

import (
	"fmt"
	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"time"
)

type DatabaseConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	Charset  string `ini:"charset"`
	Loc      string `ini:"loc"`
}
type Config struct {
	DatabaseConfig `ini:"database"`
}

var (
	DB  *gorm.DB
	err error
)

func init() {
	var configObj = new(Config)
	err = ini.MapTo(configObj, "./conf/config.ini")
	if err != nil {
		log.Fatal("config failed err:", err)
		return
	}
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		configObj.User,
		configObj.Password,
		configObj.Host,
		configObj.Port,
		configObj.Database,
		configObj.Charset,
		url.QueryEscape(configObj.Loc))
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	//数据库迁移表，第一次启动后，可以注释掉
	DB.AutoMigrate(&User{}, &Admin{}, &Article{}, &Category{})
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
