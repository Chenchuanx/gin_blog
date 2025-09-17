package handler

import (
	"errors"
	"goBlog/models"
	"goBlog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 响应成功信息
func ResponseUserSuccess(c *gin.Context, msg string, user *models.Users) {
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess, // 保持使用常量
		"msg":  msg,         // 使用传入的消息参数
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

// 登录  检查用户名和密码是否正确
func UserLogin(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	user := models.Users{
		Username: req.Username,
		Password: req.Password,
	}
	// 检查用户名和密码是否正确
	dbUser, err := service.CheckUserByPassword(GetDB(c), &user)
	if err != nil {
		code := CodeUnauthorized
		if err.Error() == "用户不存在" {
			code = CodeUserNotExists
		} else if err.Error() == "密码错误" {
			code = CodePasswordError
		}
		ResponseError(c, code, err.Error())
		return
	}
	dbUser.Password = "" // 密码不返回
	// 登录成功时
	ResponseUserSuccess(c, "登录成功", dbUser)
	// 登录日志
	log := GetLogger(c)
	log.Info("UserLogin success, userid: %d", dbUser.ID)
}

// 注册  检查用户名是否已存在、创建用户
func UserSignup(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=6,max=32"`
		Email    string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 检查用户名是否已存在
	existingUser, err := service.GetUserInfoByName(GetDB(c), req.Username)
	if err == nil && existingUser != nil { // 没有错误且查找到用户 == 用户名已存在
		ResponseError(c, CodeUserExists, "注册失败:用户名已存在")
		return
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ResponseError(c, CodeDbError, "注册失败:意外错误")
		return
	}

	// 创建用户
	newUser := models.Users{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	err = service.CreateUser(GetDB(c), &newUser)
	if err != nil {
		ResponseError(c, CodeDbError, "注册失败："+err.Error())
		return
	}
	// 注册成功时
	ResponseUserSuccess(c, "注册成功", &newUser)
	// 注册日志
	log := GetLogger(c)
	log.Info("UserSignup success, userid: %d", newUser.ID)
}

// 修改密码	检查旧密码是否正确、更新密码
func ChangePassword(c *gin.Context) {
	var req struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		NewPassword string `json:"newpassword" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 检查旧密码是否正确
	user := models.Users{
		Username: req.Username,
		Password: req.Password,
	}
	dbUser, err := service.CheckUserByPassword(GetDB(c), &user)
	if err != nil {
		code := CodeUnauthorized
		if err.Error() == "用户不存在" {
			code = CodeUserNotExists
		} else if err.Error() == "密码错误" {
			code = CodePasswordError
		}
		ResponseError(c, code, err.Error())
		return
	}

	// 更新密码
	dbUser.Password = req.NewPassword
	err = service.UpdatePassword(GetDB(c), dbUser)
	if err != nil {
		ResponseError(c, CodeDbError, "修改密码失败: "+err.Error())
		return
	}
	// 修改密码成功
	ResponseUserSuccess(c, "修改成功", dbUser)
	// 修改密码日志
	log := GetLogger(c)
	log.Info("ChangePassword success, userid: %d", dbUser.ID)
}
