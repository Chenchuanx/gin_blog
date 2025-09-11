package service

import (
	"errors"
	"goBlog/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// create User
func CreateUser(db *gorm.DB, user *models.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := db.Create(&user)
	return result.Error
}

// UpdatePassword
func UpdatePassword(db *gorm.DB, id int, user *models.Users) error {
	result := db.Select("password").Where("id = ?", id).Updates(&user)
	return result.Error
}

// get User by id
func GetUserById(db *gorm.DB, id int) (*models.Users, error) {
	var user models.Users
	result := db.Model(&user).Where("id", id).First(&user)
	return &user, result.Error
}

// get User by username
func GetUserInfoByName(db *gorm.DB, name string) (*models.Users, error) {
	var user models.Users
	result := db.Model(&user).Where("username LIKE ?", name).First(&user)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &user, result.Error
}

// check User password
func CheckUserByPassword(db *gorm.DB, user *models.Users) (*models.Users, error) {
	var dbUser models.Users

	result := db.Select("id", "password").Where("username = ?", user.Username).First(&dbUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return nil, errors.New("密码错误")
	}

	fullUser := &models.Users{}
	if err := db.First(fullUser, dbUser.ID).Error; err != nil {
		return nil, err
	}

	return fullUser, nil
}
