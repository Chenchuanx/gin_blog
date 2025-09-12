package handler

import (
	"errors"
	"goBlog/models"
	"goBlog/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 定义自定义错误码
const (
	CodeSuccess       = "0"
	CodeParamError    = "10001"
	CodeUnauthorized  = "10002"
	CodeUserExists    = "10003"
	CodeUserNotExists = "10004"
	CodeDbError       = "10005"
	CodePasswordError = "10006"
)

func UserLogin(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeParamError,
			"msg":  "参数错误：" + err.Error(),
		})
		return
	}

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
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}
	dbUser.Password = ""

	// 成功响应中包含用户信息
	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  "登录成功",
		"user": gin.H{
			"id":       dbUser.ID,
			"username": dbUser.Username,
		},
	})
}

func UserSignup(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=6,max=32"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeParamError,
			"msg":  "参数错误：" + err.Error(),
		})
		return
	}

	existingUser, err := service.GetUserInfoByName(GetDB(c), req.Username)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeUserExists,
			"msg":  "注册失败:用户名已存在",
		})
		return
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeDbError,
			"msg":  "注册失败:意外错误",
		})
		return
	}

	newUser := models.Users{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	err = service.CreateUser(GetDB(c), &newUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeDbError,
			"msg":  "注册失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  "注册成功",
		"user": gin.H{
			"id":       newUser.ID,
			"username": newUser.Username,
		},
	})
}

func ChangePassword(c *gin.Context) {
	var req struct {
		Username    string `json:"username" binding:"required"`
		Password    string `json:"password" binding:"required"`
		NewPassword string `json:"newpassword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeParamError,
			"msg":  "参数错误：" + err.Error(),
		})
		return
	}

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
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
		return
	}

	dbUser.Password = req.NewPassword
	err = service.UpdatePassword(GetDB(c), dbUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": CodeDbError,
			"msg":  "修改密码失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": CodeSuccess,
		"msg":  "修改成功",
		"user": gin.H{
			"id":       dbUser.ID,
			"username": dbUser.Username,
		},
	})
}
