package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	SUCCESS = StatusCode{0, "success"}
)

type StatusCode struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

type ResponseData struct {
	StatusCode
	Data interface{} `json:"data"`
}

func SetResponse(rep StatusCode, data interface{}) *ResponseData {
	r := new(ResponseData)
	r.Code = rep.Code
	r.Msg = rep.Msg
	r.Data = data
	return r
}

func SendResponse(ctx *gin.Context, rep StatusCode, data interface{}) {
	retData := SetResponse(rep, data)
	ctx.AbortWithStatusJSON(http.StatusOK, retData)
	return
}
