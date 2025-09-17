package service

import (
	"goBlog/models"

	"gorm.io/gorm"
)

// CreateArticle 创建文章
func CreateArticle(db *gorm.DB, article *models.Article) error {
	return db.Create(article).Error
}

// UpdateArticle 更新文章
func UpdateArticle(db *gorm.DB, article *models.Article) error {
	return db.Save(article).Error
}

// DeleteArticle 删除文章
func DeleteArticle(db *gorm.DB, id int) error {
	return db.Delete(&models.Article{}, id).Error
}

// GetArticleList 获取文章列表（使用JOIN连表查询）
func GetArticleList(db *gorm.DB) ([]*models.Article, error) {
	var articles []*models.Article

	// 预加载Users关联，并只选择需要的字段
	err := db.Model(&models.Article{}).
		Select("id", "title", "like_count", "author_id").
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username") // 只选择必要的用户字段
		}).
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(db *gorm.DB, id int) (*models.Article, error) {
	var article models.Article

	// 添加预加载Users关联，并只选择需要的字段
	err := db.Where("id = ?", id).
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username") // 选择必要的用户字段
		}).
		First(&article).Error

	if err != nil {
		return nil, err
	}
	return &article, nil
}

// 对点赞数进行增量更新
func UpdateArticleLike(db *gorm.DB, id int, count int) error {
	return db.Model(&models.Article{}).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", count)).
		Error
}

// GetMyArticles 获取当前用户的文章列表
func GetMyArticles(db *gorm.DB, authorID int) ([]*models.Article, error) {
	var articles []*models.Article

	// 预加载Users关联，并只选择需要的字段
	err := db.Model(&models.Article{}).
		Select("id", "title", "like_count", "author_id").
		Preload("Users", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username") // 只选择必要的用户字段
		}).
		Where("author_id = ?", authorID).
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
