package middleware

import (
	"goBlog/core"
	g "goBlog/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 数据库中间件 GORM
func GormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(g.CTX_DB, db)
		ctx.Next()
	}
}

// 日志中间件
func Logger(log *core.LoggerOutput) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(g.CTX_LOG, log)
		ctx.Next()
	}
}
