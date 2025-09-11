package handler

import (
	g "goBlog/global"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取 *gorm.DB
func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet(g.CTX_DB).(*gorm.DB)
}
