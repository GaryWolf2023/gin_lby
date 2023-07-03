package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
	// 定义了一个结构体，方便继承
}

func (base BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "success",
	})
}
func (base BaseController) Error404(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code": -1,
		"msg":  "404 Not Found",
	})
}
