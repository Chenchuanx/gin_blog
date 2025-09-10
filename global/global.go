package global

import (
	"goBlog/config"
)

type LoggerInterface interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}

var (
	Config *config.Config
	Logger LoggerInterface
	// Db     *gorm.DB
)
