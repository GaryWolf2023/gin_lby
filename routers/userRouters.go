package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouters() {
	// middlewares.Init这个中间件写在这个地方
	RegisterRouter(func(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {
		rgPublic.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": gin.H{},
				"msg":  "登录成功",
			})
		})
		userPrivate := rgPrivate.Group("/user")
		userPrivate.GET("", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]interface{}{
					{"id": 1, "name": "张三"},
					{"id": 2, "name": "张三"},
					{"id": 3, "name": "张三"},
					{"id": 4, "name": "张三"},
				},
			})
		})
		// 里面的函数可以抽出来放在controllers里面
		userPrivate.GET("/:id", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 200,
				"data": map[string]interface{}{
					"id":   1,
					"name": "张三",
				},
				"msg": "获取用户信息成功",
			})
		})

		userPrivate.POST("/logout", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"code": 200,
				"data": gin.H{},
				"msg":  "登出成功",
			})
		})
	})
}
