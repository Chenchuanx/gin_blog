package global

import (
	"goBlog/config"
)

// 日志接口
type LoggerInterface interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

// 上下文键
const (
	CTX_DB        = "_db_field"
	CTX_RDB       = "_rdb_field"
	CTX_USER_AUTH = "_user_auth_field"
)

// 全局变量
var (
	Config *config.Config
	Logger LoggerInterface
	// Db     *gorm.DB
)
