package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/x14n/x14n-gin-api/global"
)

const (
	SUCCEED = 0
	ERROR   = 1
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
	Time string `json:"time"`
}

func ResultJson(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
		Time: global.YYYYMMDDHHIISS,
	})
}

func OK(ctx *gin.Context) {
	ResultJson(ctx, SUCCEED, "请求成功", map[string]interface{}{})
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCEED, msg, map[string]interface{}{})
}

func OkWithData(ctx *gin.Context, msg string, data any) {
	ResultJson(ctx, SUCCEED, msg, data)
}

func Error(ctx *gin.Context) {
	ResultJson(ctx, ERROR, "请求失败", map[string]interface{}{})
}

func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, map[string]interface{}{})
}
