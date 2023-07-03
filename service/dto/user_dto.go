package dto

type UserLoginDto struct {
	UserName string `json:"username" binding:"required"` // 用户名,必填(使用binding:"required"来标识必填项)
	Password string `json:"password" binding:"required"` // 密码,必填
}
