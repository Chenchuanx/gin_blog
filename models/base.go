package models

import "time"

// 通用模型(替换gorm.Model)
// MAX_INT = 2147483647, 21亿不可能超
type Model struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
