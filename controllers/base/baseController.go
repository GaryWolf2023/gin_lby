package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (base BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "success",
	})
}
func (base BaseController) Error(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code": -1,
		"msg":  "404 Not Found",
	})
}
