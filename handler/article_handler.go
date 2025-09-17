package handler

import (
	"errors"
	"goBlog/models"
	"goBlog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 响应成功信息 - 文章列表
func ResponseArticleListSuccess(c *gin.Context, msg string, articles []*models.Article) {
	c.JSON(http.StatusOK, gin.H{
		"code":     CodeSuccess,
		"msg":      msg,
		"articles": articles,
	})
}

// 响应成功信息 - 单篇文章
func ResponseArticleSuccess(c *gin.Context, msg string, article *models.Article) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CodeSuccess,
		"msg":     msg,
		"article": article,
	})
}

// 创建文章
func CreateArticle(c *gin.Context) {
	var req struct {
		Title    string `json:"title" binding:"required,min=2,max=100"`
		Content  string `json:"content" binding:"required,min=10"`
		AuthorID int    `json:"author_id" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 创建文章对象
	article := &models.Article{
		Title:     req.Title,
		Content:   req.Content,
		AuthorID:  req.AuthorID,
		LikeCount: 0,
	}

	// 调用服务层创建文章
	if err := service.CreateArticle(GetDB(c), article); err != nil {
		ResponseError(c, CodeDbError, "创建文章失败："+err.Error())
		return
	}

	ResponseArticleSuccess(c, "创建文章成功", article)

	// 记录日志
	log := GetLogger(c)
	log.Info("CreateArticle success, article id: %d, author id: %d", article.ID, article.AuthorID)
}

// 获取文章列表
func GetArticleList(c *gin.Context) {
	articles, err := service.GetArticleList(GetDB(c))
	if err != nil {
		ResponseError(c, CodeDbError, "获取文章列表失败："+err.Error())
		return
	}

	ResponseArticleListSuccess(c, "获取文章列表成功", articles)
	// 记录日志
	log := GetLogger(c)
	log.Info("GetArticleList success, article count: %d", len(articles))
}

// 更新文章
func UpdateArticle(c *gin.Context) {
	// 从URL参数中获取文章ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseError(c, CodeParamError, "无效的文章ID")
		return
	}

	var req struct {
		Title   string `json:"title" binding:"required,min=2,max=100"`
		Content string `json:"content" binding:"required,min=10"`
	}
	if err = c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 获取原文章
	article, err := service.GetArticleDetail(GetDB(c), id)
	if err != nil {
		ResponseError(c, CodeDbError, "获取更新后的文章信息失败："+err.Error())
		return
	}

	// 更新文章字段
	article.Title = req.Title
	article.Content = req.Content

	// 调用服务层更新文章
	if err = service.UpdateArticle(GetDB(c), article); err != nil {
		ResponseError(c, CodeDbError, "更新文章失败："+err.Error())
		return
	}

	ResponseArticleSuccess(c, "更新文章成功", article)

	// 记录日志
	log := GetLogger(c)
	log.Info("UpdateArticle success, article id: %d", id)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var req struct {
		ID int `json:"author_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 调用服务层删除文章
	if err := service.DeleteArticle(GetDB(c), req.ID); err != nil {
		ResponseError(c, CodeDbError, "删除文章失败："+err.Error())
		return
	}

	// 修复结构体初始化语法
	ResponseArticleSuccess(c, "删除文章成功", &models.Article{Model: models.Model{ID: req.ID}})

	// 记录日志
	log := GetLogger(c)
	log.Info("DeleteArticle success, article id: %d", req.ID)
}

// 获取文章详情
func GetArticleDetail(c *gin.Context) {
	// 从URL参数中获取文章ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseError(c, CodeParamError, "无效的文章ID")
		return
	}

	// 调用服务层获取文章详情
	article, err := service.GetArticleDetail(GetDB(c), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseError(c, CodeUserNotExists, "文章不存在")
		} else {
			ResponseError(c, CodeDbError, "获取文章详情失败："+err.Error())
		}
		return
	}

	ResponseArticleSuccess(c, "获取文章详情成功", article)
}

func ChangeArticleLike(c *gin.Context) {
	idStr := c.Param("id")
	ID, err := strconv.Atoi(idStr)
	if err != nil {
		ResponseError(c, CodeParamError, "无效的文章ID")
		return
	}

	var req struct {
		Count int `json:"count" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	// 调用服务层更新文章点赞数
	if err := service.UpdateArticleLike(GetDB(c), ID, req.Count); err != nil {
		ResponseError(c, CodeDbError, "更新文章点赞数失败："+err.Error())
		return
	}

	ResponseArticleSuccess(c, "更新文章点赞数成功", &models.Article{Model: models.Model{ID: ID}, LikeCount: req.Count})
}

// GetMyArticles 获取当前用户的文章列表
func GetMyArticles(c *gin.Context) {
	var req struct {
		AuthorID int `json:"author_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		ResponseError(c, CodeParamError, "参数错误："+err.Error())
		return
	}

	articles, err := service.GetMyArticles(GetDB(c), req.AuthorID)
	if err != nil {
		ResponseError(c, CodeDbError, "获取我的文章失败："+err.Error())
		return
	}

	ResponseArticleListSuccess(c, "获取我的文章成功", articles)
}
