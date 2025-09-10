package core

import (
	"fmt"
	"goBlog/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() {
	global.Db = MysqlConnect()
}

func MysqlConnect() *gorm.DB {
	if global.Config.MySql.Host == "" {
		fmt.Println("不存在")
	}
	dsn := global.Config.MySql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	return db
}
