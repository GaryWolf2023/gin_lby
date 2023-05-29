package routers

import (
	"Gin_one/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.Default()
	// 先加载路由信息然后在启动服务
	r.Use(middlewares.Cors())
	UserRouters(r)
	RoleRouters(r)
	DefaultRouters(r)
	r.Run(":10010")
}
