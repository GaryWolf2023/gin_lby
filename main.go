package main

import (
	"Gin_one/config"
	"Gin_one/routers"
)

func main() {
	config.InitConfig()
	// 初始化路由并创建一个网络服务
	routers.InitRouters()
	// r.Run(":10001") // 启动一个web服务
}
