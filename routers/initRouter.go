package routers

import (
	"Gin_one/config"
	docs "Gin_one/docs"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// 初始化路由 并创建一个网络服务
func InitRouters() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	// 先加载路由信息然后在启动服务
	rgPublic := r.Group("/api/v1/public")
	rgPrivate := r.Group("/api/v1")

	// 调用这个方法让所有的路由都存储在gFnRouter中
	InitBasePlatformRouter(r)

	// 加载所有的路由将所有的路由注册到路由组（rgPublic，rgPrivate）中
	for _, registerRouter := range gFnRouter {
		registerRouter(rgPublic, rgPrivate)
	}
	// 从配置文件中读取端口号
	serverPort := config.ViperConfig.Get("server.port")
	if serverPort == "" {
		serverPort = "8080"
	}
	// 从配置文件中读取是否开启swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 启动服务
	// serverError := r.Run(fmt.Sprintf(":%s", serverPort))
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: r,
	}
	go func() {
		// 创建了一个携程监听server服务
		// if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// 	log.Fatalf("listen: %s\n", err)
		// 	return
		// }
		if err := server.ListenAndServe(); err != nil && err != server.Shutdown(ctx) {
			log.Fatalf("listen: %s\n", err)
			return
		}
	}()

	<-ctx.Done()

	// if serverError != nil {
	// 	panic(fmt.Sprintf("启动服务失败%s", serverError.Error()))
	// } else {
	// 	fmt.Println("启动服务成功 端口:", serverPort)
	// }
	stop()
	// TODO: 记录日志---服务停止成功
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// 创建一个5秒的超时上下文，超时后会强制关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		// TODO: 记录日志---服务停止失败
		log.Fatal("Server forced to shutdown: ", err)
	}
	// TODO: 记录日志---服务停止成功
	log.Println("Server exiting")
}

// 初始化所有的路由
func InitBasePlatformRouter(r *gin.Engine) {
	// r.Use(middlewares.Cors())
	InitUserRouters()
}
