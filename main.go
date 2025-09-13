package main

import (
	"goBlog/core"
	"goBlog/middleware"
	"goBlog/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	core.InitConf()
	log := core.InitLogger()
	db := core.InitGorm()

	gin.SetMode("release")
	r := gin.New()
	r.Use(middleware.CORS()) // 添加CORS中间件，确保在所有路由之前
	r.Use(middleware.GormDB(db))
	r.Use(middleware.Logger(log))
	routers.InitRouter(r)
	// fmt.Println("Server Start ...")
	log.Info("Server Start ...")
	r.Run("127.0.0.1:8080")

	// fmt.Println(global.Config)
	// global.Logger.Debug("这是 debug 日志") // 如果 level >= info，则不输出
	// global.Logger.Info("这是 info 日志")
	// global.Logger.Warning("这是 warning 日志")
	// global.Logger.Error("这是 error 日志")
	// global.Logger.Fatal("这是 fatal 日志")
}
