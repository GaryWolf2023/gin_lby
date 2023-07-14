package dto

type UserLoginDto struct {
	UserName string `json:"username" binding:"required,email,first_name_a" message:"用户名必须以a开头" required_err:"用户名不能为空"` // 用户名,必填(使用binding:"required"来标识必填项)
	Password string `json:"password" binding:"required" message:"密码不能为空"`                                              // 密码,必填
}
