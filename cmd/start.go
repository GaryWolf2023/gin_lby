package cmd

import (
	"Gin_one/config"
	"Gin_one/global"
	"Gin_one/routers"
	"Gin_one/utils"
)

// 服务启动 读取配置文件 初始化数据库 初始化路由并创建一个网络服务
func Start() {
	var initError error
	// 有先后顺序，我们必须先初始化配置文件最后在创建网络服务
	config.InitConfig() // 初始化配置文件

	global.Logger = config.InitLogger() // 初始化日志

	db, err := config.InitDB() // 初始化数据库
	global.DB = db
	if err != nil {
		initError = utils.AppendError(initError, err)
	}

	rdb, err := config.InitRedis() // 初始化redis
	global.Rdb = rdb
	if err != nil {
		initError = utils.AppendError(initError, err)
	}
	_ = global.Rdb.Set("test", "test")
	// 将错误信息打印到日志中
	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}
	}
	// sql.InitSql()
	// 初始化数据库
	// 最后初始化路由服务
	routers.InitRouters() // 初始化路由并创建一个网络服务
}
