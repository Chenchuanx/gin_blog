package routers

import (
	"goBlog/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/", handler.UserLogin)
	r.POST("/sign_up", handler.UserSignup)
}
