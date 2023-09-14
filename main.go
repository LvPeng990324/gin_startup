package main

import (
	"contract_system_server/common"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	common.InitConfig()
	// 初始化日志
	common.InitLog()
	// 初始化数据库
	db := common.InitDB()
	defer db.Close()
	// 初始化redis
	rdc := common.InitRedis()
	defer rdc.Close()

	// 设置启动模式
	debug_flag := viper.GetBool("server.debug")
	if debug_flag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	r := gin.Default()
	r = CollectRoute(r)
	// 运行
	port := viper.GetString("server.port")

	common.AddInfo("server running on port:" + port)

	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
