package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// 通用模型(替换gorm.Model)
// MAX_INT = 2147483647, 21亿不可能超
type Model struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 用户模型
type Users struct {
	Model
	Username string `json:"username" gorm:"unique;type:varchar(30)not null"`
	Password string `json:"password" gorm:"type:varchar(50);not null"`
	Email    string `json:"email" gorm:"type:varchar(30);not null"`
	Role     int    `json:"role" gorm:"type:int;DEFAULT:2;not null"`
}

func GetUserById(db *gorm.DB, id int) (*Users, error) {
	var user Users
	result := db.Model(&user).Where("id", id).First(&user)
	return &user, result.Error
}

func GetUserInfoByName(db *gorm.DB, name string) (*Users, error) {
	var user Users
	result := db.Model(&user).Where("username LIKE ?", name).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, result.Error
}

// func CreateUser(db)
