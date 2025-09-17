package routers

import (
	"goBlog/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// 处理前端请求接口
	// 用户相关路由
	r.POST("/login", handler.UserLogin)
	r.POST("/sign_up", handler.UserSignup)
	r.POST("/change_password", handler.ChangePassword)

	// 文章相关路由
	r.POST("/articles/create", handler.CreateArticle)
	r.GET("/articles/get", handler.GetArticleList)
	r.POST("/articles/getmy", handler.GetMyArticles)
	r.GET("/articles/get/:id", handler.GetArticleDetail)
	r.POST("/articles/update/:id", handler.UpdateArticle)
	r.POST("/articles/delete/:id", handler.DeleteArticle)
	r.POST("/articles/like/:id", handler.ChangeArticleLike)

}
