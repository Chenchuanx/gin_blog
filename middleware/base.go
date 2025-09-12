package middleware

import (
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
