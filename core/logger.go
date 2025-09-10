package core

import (
	"fmt"
	"goBlog/global"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ========== 日志级别定义 ==========
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelMap = map[string]LogLevel{
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
	"fatal": FATAL,
}

// LoggerOutput 自定义日志器
type LoggerOutput struct {
	logger *log.Logger
	level  LogLevel
	mu     sync.Mutex
}

var (
	instance *LoggerOutput
	logFile  *os.File
)

// InitLogger 初始化日志系统
func InitLogger() {
	cfg := global.Config.Logger

	// 解析级别
	lvl, exists := levelMap[strings.ToLower(cfg.Level)]
	if !exists {
		lvl = INFO // 默认 info
	}

	// 创建目录
	if err := os.MkdirAll(cfg.Director, 0755); err != nil {
		panic("无法创建日志目录: " + err.Error())
	}

	// 打开文件
	logPath := filepath.Join(cfg.Director, "app.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("无法打开日志文件: " + err.Error())
	}
	logFile = file

	// 输出目标
	var writer io.Writer = file
	if cfg.LogInConsole {
		writer = io.MultiWriter(file, os.Stdout)
	}

	// 前缀
	prefix := cfg.Prefix
	if prefix == "" {
		prefix = "[LOG] "
	} else if !hasSuffix(prefix, " ") {
		prefix += " "
	}

	// 创建标准日志器
	stdLogger := log.New(writer, prefix, log.Ldate|log.Ltime|log.Lmicroseconds)

	// 创建封装日志器
	instance = &LoggerOutput{
		logger: stdLogger,
		level:  lvl,
	}

	// 赋值给全局变量
	global.Logger = instance
}

// CloseLogger 关闭日志文件
func CloseLogger() {
	if logFile != nil {
		_ = logFile.Close()
		logFile = nil
	}
}

// ========== 实现日志方法 ==========
func (l *LoggerOutput) Debug(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if DEBUG >= l.level {
		l.logger.Println("[DEBUG] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Info(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if INFO >= l.level {
		l.logger.Println("[INFO ] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Warning(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if WARN >= l.level {
		l.logger.Println("[WARN ] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Error(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if ERROR >= l.level {
		l.logger.Println("[ERROR] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Fatal(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if ERROR >= l.level {
		l.logger.Println("[FATAL] " + fmt.Sprintf(format, v...))
	}
}

// 工具函数
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
