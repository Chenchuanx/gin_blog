package models

// 用户模型
type Users struct {
	Model
	Username string `json:"username" gorm:"unique;type:varchar(30);not null"` // 修复标签格式，添加分号
	Password string `json:"password" gorm:"type:varchar(64);not null"`
	Email    string `json:"email" gorm:"type:varchar(30);not null"`
	Role     int    `json:"role" gorm:"type:int;DEFAULT:2;not null"`
}
