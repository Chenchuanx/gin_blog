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
	"time"
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
	logger       *log.Logger
	level        LogLevel
	mu           sync.Mutex // 写日志锁, 避免并发写日志
	logFile      *os.File   // 当前日志文件
	currentDate  string
	logDirectory string // 日志目录
	logInConsole bool   // 是否同时输出到控制台
	prefix       string
}

var (
	instance *LoggerOutput
)

// InitLogger 初始化日志系统
func InitLogger() *LoggerOutput {
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

	// 获取当前日期
	currentDate := time.Now().Format("2006-01-02")

	// 打开今天的日志文件
	logPath := filepath.Join(cfg.Director, fmt.Sprintf("log_%s.log", currentDate))
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic("无法打开日志文件: " + err.Error())
	}

	// 设置输出目标
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
		logger:       stdLogger,
		level:        lvl,
		logFile:      file,
		currentDate:  currentDate,
		logDirectory: cfg.Director,
		logInConsole: cfg.LogInConsole,
		prefix:       prefix,
	}

	return instance
}

// CloseLogger 关闭日志文件
func CloseLogger() {
	if instance != nil && instance.logFile != nil {
		_ = instance.logFile.Close()
		instance.logFile = nil
	}
}

// 检查是否需要切换日志文件
func (l *LoggerOutput) checkAndRotateLogFile() {
	// 获取当前日期
	today := time.Now().Format("2006-01-02")

	// 如果日期变更，创建新的日志文件
	if today != l.currentDate {
		// 关闭旧文件
		if l.logFile != nil {
			_ = l.logFile.Close()
		}

		// 打开新文件
		logPath := filepath.Join(l.logDirectory, fmt.Sprintf("app_%s.log", today))
		var err error
		l.logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			// 错误处理：输出到标准错误
			fmt.Fprintf(os.Stderr, "无法打开新的日志文件: %v\n", err)
			return
		}

		// 更新输出目标
		var writer io.Writer = l.logFile
		if l.logInConsole {
			writer = io.MultiWriter(l.logFile, os.Stdout)
		}

		// 更新日志器
		l.logger = log.New(writer, l.prefix, log.Ldate|log.Ltime|log.Lmicroseconds)
		l.currentDate = today
	}
}

// ========== 实现日志方法 ==========
func (l *LoggerOutput) Debug(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if DEBUG >= l.level {
		l.checkAndRotateLogFile() // 检查是否需要切换日志文件
		l.logger.Println("[DEBUG] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Info(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if INFO >= l.level {
		l.checkAndRotateLogFile()
		l.logger.Println("[INFO] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Warning(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if WARN >= l.level {
		l.checkAndRotateLogFile()
		l.logger.Println("[WARN] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Error(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if ERROR >= l.level {
		l.checkAndRotateLogFile()
		l.logger.Println("[ERROR] " + fmt.Sprintf(format, v...))
	}
}

func (l *LoggerOutput) Fatal(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if ERROR >= l.level {
		l.checkAndRotateLogFile()
		l.logger.Println("[FATAL] " + fmt.Sprintf(format, v...))
	}
}

// 工具函数
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
