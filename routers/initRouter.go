package routers

import (
	"Gin_one/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 这里是定义的一个函数类型，这个函数类型的参数是两个路由组，一个是公共的路由组，一个权限路由组（token，role之类的）
type IFnRegisterRouter = func(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup)

var gFnRouter []IFnRegisterRouter

func RegisterRouter(fn IFnRegisterRouter) {
	if fn == nil {
		return
	}
	gFnRouter = append(gFnRouter, fn)
}

func InitRouters() {
	r := gin.Default()
	// 先加载路由信息然后在启动服务
	rgPublic := r.Group("/api/v1/public")
	rgPrivate := r.Group("/api/v1")

	// 调用这个方法让所有的路由都存储在gFnRouter中
	InitBasePlatformRouter(r)

	// 加载所有的路由将所有的路由注册到路由组（rgPublic，rgPrivate）中
	for _, registerRouter := range gFnRouter {
		registerRouter(rgPublic, rgPrivate)
	}

	serverPort := config.ViperConfig.Get("server.port")
	fmt.Println("serverPort", serverPort)
	if serverPort == "" {
		serverPort = "8080"
	}

	serverError := r.Run(fmt.Sprintf(":%s", serverPort))

	if serverError != nil {
		panic(fmt.Sprintf("启动服务失败%s", serverError.Error()))
	} else {
		fmt.Println("启动服务成功 端口:", serverPort)
	}
}

// 初始化所有的路由
func InitBasePlatformRouter(r *gin.Engine) {
	// r.Use(middlewares.Cors())
	InitUserRouters()
}
