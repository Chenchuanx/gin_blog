package handler

import (
	g "goBlog/global"

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

// 获取 中间件 *gorm.DB
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}
