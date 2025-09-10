package main

import (
	"goBlog/core"
	"goBlog/global"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.InitGorm()

	// fmt.Println(global.Config)
	global.Logger.Debug("这是 debug 日志") // 如果 level >= info，则不输出
	global.Logger.Info("这是 info 日志")
	global.Logger.Warning("这是 warning 日志")
	global.Logger.Error("这是 error 日志")
	global.Logger.Fatal("这是 fatal 日志")
}
