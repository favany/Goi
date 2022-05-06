package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1001, // 程序中的错误码
	"message": ""， // 提示信息
	"data": {},    // 数据
}
*/

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
