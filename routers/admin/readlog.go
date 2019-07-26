package admin

import (
	"server/models"
	"server/util"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type duration struct {
	Begin int64 `json:"begin" binding:"required"`
	End   int64 `json:"end" binding:"required"`
}

//ReadLogCount 阅读记录计数
func ReadLogCount(ctx *gin.Context) {
	param := duration{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	begin, end := time.Unix(param.Begin, 0), time.Unix(param.End, 0)
	count := 0
	models.Db().Model(&models.ReadLog{}).Where("time BETWEEN ? AND ?", begin, end).Count(&count)
	util.SetResponse(ctx, count)
}

//ReadLogs 阅读记录
func ReadLogs(ctx *gin.Context) {
	param := duration{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	begin, end := time.Unix(param.Begin, 0), time.Unix(param.End, 0)
	logs := []models.ReadLog{}
	models.Db().Where("time BETWEEN ? AND ?", begin, end).Order("time").Find(&logs)
	util.SetResponse(ctx, &logs)
}
