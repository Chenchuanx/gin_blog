package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	CTX_DB        = "_db_field"
	CTX_RDB       = "_rdb_field"
	CTX_USER_AUTH = "_user_auth_field"
)

func WithGormDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(CTX_DB, db)
		ctx.Next()
	}
}
