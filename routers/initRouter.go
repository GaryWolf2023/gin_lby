package routers

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()
	// 先加载路由信息然后在启动服务
	UserRouters(r)
	RoleRouters(r)
	DefaultRouters(r)
	r.Run(":8888")
}
