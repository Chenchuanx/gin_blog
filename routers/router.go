package routers

import (
	"goBlog/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/login", handler.UserLogin)
	r.POST("/sign_up", handler.UserSignup)
	r.POST("/change_password", handler.ChangePassword)
}
