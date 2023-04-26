package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func DefaultRouters(r *gin.Engine) {
	// 基本上所有的路由必须使用middleware.Init()中间件来判断登陆状态
	// 定义一个路由组使用Group
	r.GET("/", func(c *gin.Context) {
		fmt.Println("欢迎来到gin")
	}, func(ctx *gin.Context) {
		// ctx.String(200, "欢迎来到gin")
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "欢迎来到gin",
		})
	})
	// http://localhost:8888/test_jsonp?callback=xxx
	// 发起一个jsonp请求
	// 解决跨域请求
	r.GET("test_jsonp", func(ctx *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		ctx.JSONP(200, data)
	})
}
