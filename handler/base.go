package handler

import (
	"goBlog/core"
	g "goBlog/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 自定义错误码
const (
	// 成功
	CodeSuccess = "0"
	// 参数错误
	CodeParamError = "10001"
	// 未授权
	CodeUnauthorized = "10002"
	// 用户已存在
	CodeUserExists = "10003"
	// 用户不存在
	CodeUserNotExists = "10004"
	// 数据库错误
	CodeDbError = "10005"
	// 密码错误
	CodePasswordError = "10006"
)

// 响应错误信息
func ResponseError(c *gin.Context, code string, errMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errMsg,
	})
}

// 获取 中间件 *gorm.DB
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}

func GetLogger(c *gin.Context) *core.LoggerOutput {
	return c.MustGet(g.CTX_LOG).(*core.LoggerOutput)
}
