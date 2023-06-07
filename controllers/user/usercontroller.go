package user

import (
	"Gin_one/controllers/base"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

func NewUserController() UserController {
	return UserController{}
}

// @Summary 用户登录
// @Description 用户登录
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {string} json "{"code":200,"data":{"token":},"msg":"登录成功"}"
// @Failure 401 {string} json "{"code":400,"msg":"登录失败"}"
// @Router /api/v1/public/user/login [post]
func (usercon UserController) UserLogin(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(200, gin.H{
		"code": 200,
		"data": "1231293g2y3vbf19763d2yn3d1y73bfg719cy7bdu81mch28u",
		"msg":  "登录成功",
	})
}

// 抽离控制器
// 改造成一个结构体，方便继承
// 我们写一个通用的控制器，比如成功失败，然后我们将它继承过来，就可以使用他里面的方法了
func (usercon UserController) UserRoute(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	fmt.Println(username)
	v, ok := username.(string) // 类型断言
	if ok {
		ctx.String(200, "用户----"+v)
	} else {
		ctx.String(200, "用户路由")
	}
}

func (usercon UserController) UserLogout(ctx *gin.Context) {
	ctx.String(200, "用户登出")
}
func (usercon UserController) UserInfo(ctx *gin.Context) {
	ctx.String(200, "用户详情")
}
func (usercon UserController) UserResetPassword(ctx *gin.Context) {
	// ctx.String(200, "更改密码")
	usercon.Success(ctx, nil)
}
