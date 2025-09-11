package main

import (
	"fmt"
	"goBlog/core"
	"goBlog/middleware"
	"goBlog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	core.InitConf()
	core.InitLogger()
	db := core.InitGorm()

	gin.SetMode("release")
	r := gin.New()
	r.Use(middleware.GormDB(db))
	routers.InitRouter(r)
	fmt.Println("Server Start ...")
	r.Run("127.0.0.1:8080")
	// fmt.Println(global.Config)
	// global.Logger.Debug("这是 debug 日志") // 如果 level >= info，则不输出
	// global.Logger.Info("这是 info 日志")
	// global.Logger.Warning("这是 warning 日志")
	// global.Logger.Error("这是 error 日志")
	// global.Logger.Fatal("这是 fatal 日志")
}
