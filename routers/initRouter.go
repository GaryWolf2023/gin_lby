package routers

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()
	UserRouters(r)
	RoleRouters(r)
	DefaultRouters(r)
	r.Run(":8888")
}
