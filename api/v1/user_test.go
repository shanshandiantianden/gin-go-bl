package v1

import (
	"fmt"
	"gin-go-bl/internal/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"os"
	"testing"
	"time"
)

var DB *gorm.DB

func init() {
	var err error
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		"realy",
		"Chen1224",
		"rm-bp16v29co3893zkyplo.mysql.rds.aliyuncs.com",
		"3306",
		"gin-bl",
		"utf8",
		url.QueryEscape("Asia/Shanghai"))
	// 连接数据库

	NewLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢查询阈值
		LogLevel:      logger.Info, //log lever
		Colorful:      true,        //禁用彩色打印
	})

	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
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

}
func TestDelUser(t *testing.T) {
	u := services.NewUserService(DB)
	a, b := u.Delete("09616073-c485-4e67-9c74-f22f93bf2502")
	fmt.Println(a, b)
}
