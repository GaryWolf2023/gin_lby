package routers

import (
	"Gin_one/controllers/user"
	"Gin_one/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.Engine) {
	// middlewares.Init这个中间件写在这个地方
	userRouters := r.Group("/user", middlewares.Init)
	{
		userRouters.GET("/", user.UserController{}.UserRoute)
		userRouters.GET("/login", user.UserController{}.UserLogin)
		userRouters.GET("/logout", user.UserController{}.UserLogout)
		userRouters.GET("/info", user.UserController{}.UserInfo)
		userRouters.GET("/resetpassword", user.UserController{}.UserResetPassword)
	}
}
