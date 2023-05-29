package main

import (
	"Gin_one/config"
	"Gin_one/routers"
	"Gin_one/sql"
)

func main() {
	config.InitConfig()
	sql.InitSql()
	// 初始化路由并创建一个网络服务
	routers.InitRouters()
}
