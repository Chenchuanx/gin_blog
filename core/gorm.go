package core

import (
	"goBlog/global"
	"goBlog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化ORM, 连接数据库
func InitGorm() *gorm.DB {
	dsn := global.Config.MySql.Dsn()                      // 数据库连接字符串
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // 连接MySQL数据库
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 迁移数据库
	if err := db.AutoMigrate((&models.Users{})); err != nil {
		panic("failed to migrate database" + err.Error())
	}
	if err := db.AutoMigrate((&models.Article{})); err != nil {
		panic("failed to migrate database" + err.Error())
	}

	return db
}
