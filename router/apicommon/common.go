package apicommon

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Error  *Error      `json:"error,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

type Error struct {
	Code int64  `json:"code"`
	Msg  string `json:"message"`
	Data string `json:"data,omitempty"`
}

type SuccessResponse struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

// ReturnSuccessResponse 接口返回数据信息
func ReturnSuccessResponse(ctx *gin.Context, code int64, msg string, data interface{}) {
	response := &SuccessResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
	return
}

// ReturnErrorResponse 返回错误
func ReturnErrorResponse(ctx *gin.Context, code int64, msg interface{}, data interface{}) {
	response := &ErrorResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	ctx.JSON(http.StatusOK, response)
	return
}
