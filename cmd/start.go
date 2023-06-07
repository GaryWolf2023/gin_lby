package cmd

import (
	"Gin_one/config"
	"Gin_one/routers"
	"Gin_one/sql"
)

// 服务启动 读取配置文件 初始化数据库 初始化路由并创建一个网络服务
func Start() {
	config.InitConfig()
	sql.InitSql()
	// 初始化路由并创建一个网络服务
	routers.InitRouters()
}
