package models

// Article 文章模型
// 与Users表通过AuthorID建立外键关联
type Article struct {
	Model
	Title     string `gorm:"type:varchar(100);not null" json:"title"`
	Content   string `gorm:"type:longtext;not null" json:"content"`
	LikeCount int    `gorm:"type:int;default:0" json:"like_count"`

	Users    Users `gorm:"foreignKey:AuthorID;references:ID; onDelete:CASCADE"`
	AuthorID int   `gorm:"type:int;not null;index"`
}
