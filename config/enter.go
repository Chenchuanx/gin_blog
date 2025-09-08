package config

type Config struct {
	MySql  MySql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}

type MySql struct {
	Host     string `yaml:"host"`
	Post     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"` // 日志等级
}

type Logger struct {
	Level        string `yaml:"level"`
	Prefix       int    `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     string `yaml:"show_line"`
	LogInConsole string `yaml:"log_in_console"`
}

type System struct {
	Host string `yaml:"host"`
	Post int    `yaml:"port"`
	Env  string `yaml:"env"`
}
