package routers

import (
	"Gin_one/middlewares"
	"Gin_one/service"

	"github.com/gin-gonic/gin"
)

func RoleRouters(r *gin.Engine) {
	roleRouters := r.Group("/role", middlewares.Init)
	{
		roleRouters.GET("/api/list", service.GetRoleList)
		roleRouters.POST("/api/create", service.CreateRole)
		roleRouters.PUT("/api/modify", service.ModifyRole)
		roleRouters.DELETE("/api/delete", service.DeleteRole)
		roleRouters.GET("/api/search", service.SearchRole)
		roleRouters.GET("/menu/list", func(ctx *gin.Context) {
			ctx.String(200, "菜单权限列表")
		})
		roleRouters.POST("/menu/create", func(ctx *gin.Context) {
			ctx.String(200, "创建菜单权限")
		})
		roleRouters.PUT("/menu/modify", func(ctx *gin.Context) {
			ctx.String(200, "修改菜单权限")
		})
		roleRouters.DELETE("/menu/delete", func(ctx *gin.Context) {
			ctx.String(200, "删除菜单权限")
		})
		roleRouters.GET("/menu/search", func(ctx *gin.Context) {
			ctx.String(200, "查询修改权限")
		})
	}
}
