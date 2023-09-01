package conf

import (
	"gin-go-bl/coveralls/models"
	"github.com/go-ini/ini"
	"log"
)

var ConfigObj = new(models.Config)

func ConfInit() {
	//读取配置文件
	err := ini.MapTo(ConfigObj, "./conf/config.ini")
	if err != nil {
		log.Fatal("config failed err:", err)
		return
	}
}
