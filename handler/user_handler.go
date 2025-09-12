package handler

import (
	"goBlog/models"
	"goBlog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	//
	var req struct {
		Username string `json:"username" binding:"required"` // 参数校验（Handler 职责）
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误：" + err.Error()})
		return
	}

	// 调用 Service 层处理业务（不关心业务细节，只传参、收结果）
	user := models.Users{
		Username: req.Username,
		Password: req.Password,
	}
	dbUser, err := service.CheckUserByPassword(GetDB(c), &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	dbUser.Password = ""

	// 3. 封装 HTTP 响应
	c.JSON(http.StatusOK, gin.H{"User": dbUser, "msg": "登录成功"})
}

func UserSignup(c *gin.Context) {
	// 注册请求参数结构体
	var req struct {
		Username string `json:"username" binding:"required,min=3,max=20"` // 用户名
		Password string `json:"password" binding:"required,min=6,max=32"` // 密码
		Email    string `json:"email" binding:"required,email"`           // 邮箱：需符合邮箱格式
	}

	// 解析并校验请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误：" + err.Error(), // 返回具体的参数错误信息（如格式不符）
		})
		return
	}

	newUser := models.Users{
		Username: req.Username,
		Password: req.Password, // 在service 加密
		Email:    req.Email,
	}

	// 调用Service层创建用户
	err := service.CreateUser(GetDB(c), &newUser)
	if err != nil {
		// 根据业务错误类型返回对应提示（如用户名已存在）
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "注册失败：" + err.Error(),
		})
		return
	}

	// 返回用户ID
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
		"user": gin.H{
			"id": newUser.ID,
		},
	})
}

func ChangePassword(c *gin.Context) {
	// 修改密码参数结构体	需要通过验证？暂时不用
	var req struct {
		Username    string `json:"username" binding:"required"` // 参数校验（Handler 职责）
		Password    string `json:"password" binding:"required"`
		NewPassword string `json:"newpassword" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "参数错误：" + err.Error(), // 返回具体的参数错误信息（如格式不符）
		})
		return
	}

	user := models.Users{
		Username: req.Username,
		Password: req.Password,
	}

	// 先检查 用户存在/原密码
	dbUser, err := service.CheckUserByPassword(GetDB(c), &user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 修改密码
	dbUser.Password = req.NewPassword
	service.UpdatePassword(GetDB(c), dbUser)

	// 返回用户ID
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"user": gin.H{
			"id":       dbUser.ID,
			"username": dbUser.Username,
		},
	})
}
