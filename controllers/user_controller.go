package controllers

import (
	"Gin_one/service/dto"
	"Gin_one/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	BaseController
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
	var userLogin dto.UserLoginDto // 定义一个dto的类型，用来接收参数
	// 在ctx中绑定参数到userLogin中，如果绑定失败，返回错误信息
	errs := ctx.ShouldBind(&userLogin)
	if errs != nil {
		Fail(ctx, ResponseJson{
			Msg: utils.ProcessValidataErrors(errs.(validator.ValidationErrors), &userLogin).Error(),
		})
		return
	}
	OK(ctx, ResponseJson{
		Data: userLogin,
	})
}

func (usercon UserController) UserLogout(ctx *gin.Context) {
	ctx.String(200, "用户登出")
}
