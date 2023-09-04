package test

import (
	"fmt"
	"gin-go-bl/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"testing"
	"time"
)

func TestDel(t *testing.T) {
	MysqlServicesInit()
	var uuid uuid.UUID
	err := uuid.UnmarshalText([]byte("f91a9b60-1d2f-42a9-b7a5-0e58a5adbac2"))
	if err != nil {
		return
	}
	fmt.Println(uuid)
	//fmt.Println(u.Value())
	fmt.Println(DeleteUser2asdaac(uuid))
}

var DB *gorm.DB
var err error

func MysqlServicesInit() {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		"realy",
		"Chen1224",
		"rm-bp16v29co3893zkyplo.mysql.rds.aliyuncs.com",
		"3306",
		"gin-bl",
		"utf8",
		url.QueryEscape("Asia/Shanghai"))
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to open database: " + err.Error())
	}

	////数据库迁移表，第一次启动后，可以注释掉
	//err := DB.AutoMigrate(&Models.User{}, &Models.Admin{}, &Models.Article{}, &Models.Category{})
	//if err != nil {
	//
	//	return
	//}
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

}
func DeleteUser2asdaac(uuid uuid.UUID) int {
	fmt.Println(uuid)
	//var user Models.User
	//err := DB.Where("uuid = ?", uuid).Delete(&user).Error
	err = DB.Raw("delete from user where uuid = ?", uuid).Error
	if err != nil {
		fmt.Println(err, 1111)
		recover()
		return utils.ERROR
	}
	return utils.SUCCESS

}
