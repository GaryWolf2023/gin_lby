package main

import (
	"Gin_one/config"
	"Gin_one/routers"
	"fmt"
)

func main() {
	fmt.Println("Hello dj!")
	config.InitConfig()
	// 初始化路由并创建一个网络服务
	routers.InitRouters()
	// r.Run(":10001") // 启动一个web服务
}
