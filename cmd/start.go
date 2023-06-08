package cmd

import (
	"Gin_one/config"
	"Gin_one/global"
	"Gin_one/routers"
	"Gin_one/sql"
)

// 服务启动 读取配置文件 初始化数据库 初始化路由并创建一个网络服务
func Start() {
	// 有先后顺序，我们必须先初始化配置文件最后在创建网络服务
	config.InitConfig()                 // 初始化配置文件
	global.Logger = config.InitLogger() // 初始化配置文件
	sql.InitSql()                       // 初始化数据库
	routers.InitRouters()               // 初始化路由并创建一个网络服务
}
