package controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int         `json:"-"`
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (m ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseJson{})
}

func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func BaseStatusReturn(resp ResponseJson, notDefailtStatus int) int {
	if resp.Status == 0 {
		return notDefailtStatus
	}
	return resp.Status
}

func OK(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, BaseStatusReturn(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, BaseStatusReturn(resp, http.StatusBadRequest), resp)
}

func ServerError(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, BaseStatusReturn(resp, http.StatusInternalServerError), resp)
}
