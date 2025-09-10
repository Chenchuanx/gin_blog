package core

import (
	"fmt"
	"goBlog/global"
	"goBlog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化ORM
func InitGorm() *gorm.DB {
	if global.Config.MySql.Host == "" {
		fmt.Println("不存在")
	}
	dsn := global.Config.MySql.Dsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	db.AutoMigrate((&models.Users{}))

	return db
}
