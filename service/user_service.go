package service

import (
	"errors"
	"goBlog/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// create User
func CreateUser(db *gorm.DB, user *models.Users) error {
	// bcrypt 加密密码， 不可逆的单向哈希算法
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败：" + err.Error())
	}
	user.Password = string(hashedPassword)
	result := db.Create(&user)
	if result.Error != nil {
		return errors.New("创建用户失败:" + result.Error.Error())
	}
	return nil
}

// UpdatePassword
func UpdatePassword(db *gorm.DB, user *models.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New("密码加密失败：" + err.Error())
	}

	user.Password = string(hashedPassword)
	result := db.Select("password").Where("id = ?", user.ID).Updates(&user)
	if result.Error != nil {
		return errors.New("数据库更新失败：" + result.Error.Error())
	}

	// 检查是否有记录被更新（避免更新不存在的用户）
	if result.RowsAffected == 0 {
		return errors.New("未找到用户，更新失败")
	}

	return nil
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
	result := db.Model(&user).Where("username = ?", name).First(&user) // 使用精确匹配替代LIKE
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// check User password;
// return user.ID, user.Password
func CheckUserByPassword(db *gorm.DB, user *models.Users) (*models.Users, error) {
	var dbUser models.Users

	// 查找用户, 找不到即
	result := db.Select("id", "password").Where("username = ?", user.Username).First(&dbUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, result.Error
	}

	// 验证密码, 密码正确返回 nil
	err := bcrypt.CompareHashAndPassword(
		[]byte(dbUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		return nil, errors.New("密码错误")
	}

	return &dbUser, nil
}
