package config

import (
	"fmt"
	"strconv"
)

type MySql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级
}

// Dsn 数据库连接字符串
func (m MySql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.User, m.Password, m.Host, strconv.Itoa(m.Port), m.DB)
	// return "root:123456@tcp(127.0.0.1:3306)/gormTest?charset=utf8mb4&parseTime=True&loc=Local"
}
