package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// errcode 从上到下 0 1 2 3 ...
const (
	OK = iota
	ParamError
	UserNotFound
	ArticleCreateFailed
	ArticleNotFound
	CreateReadLogFailed
	AlreadyPunched
	CodeConvertFailed
	GenerateUUIDFailed
	ReadCodeResponseFailed
	CreateUserFailed
	UpdateUserFailed
)

//SetError 设置错误
func SetError(ctx *gin.Context, code int, err string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err,
	})
}

//SetResponse 设置响应格式
func SetResponse(ctx *gin.Context, jsonObj interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": OK,
		"msg":  "okk",
		"data": jsonObj,
	})
}
