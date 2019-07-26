package admin

import (
	"server/models"
	"server/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type (
	idParam struct {
		ID uint `json:"ID" binding:"required"`
	}
	userResponse struct {
		ID           uint             `json:"ID"`
		Name         string           `json:"Name"`
		SchoolID     string           `json:"SchoolID"`
		RegisterTime time.Time        `json:"RegisterTime"`
		Logs         []models.ReadLog `json:"Readlogs"`
	}
)

//UserInfo 用户信息
func UserInfo(ctx *gin.Context) {
	param := idParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	var user models.User
	models.Db().First(&user, "id = ?", param.ID)
	if user.ID == 0 {
		util.SetError(ctx, util.UserNotFound, "no such user")
		return
	}

	logs := []models.ReadLog{}
	models.Db().Where("user_id = ?", user.ID).Order("time desc").Find(&logs)

	util.SetResponse(ctx, &userResponse{ID: user.ID, Name: user.Name, SchoolID: user.SchoolID, RegisterTime: user.CreatedAt, Logs: logs})
}

//UserCount 用户数量
func UserCount(ctx *gin.Context) {
	count := 0
	models.Db().Model(&models.User{}).Count(&count)
	util.SetResponse(ctx, count)
}

//AllUser 所有用户
func AllUser(ctx *gin.Context) {
	users := []models.User{}
	models.Db().Order("created_at desc").Find(&users)
	rsp := []userResponse{}
	for _, user := range users {
		logs := []models.ReadLog{}
		models.Db().Where("user_id = ?", user.ID).Order("time desc").Find(&logs)
		rsp = append(rsp, userResponse{ID: user.ID, Name: user.Name, SchoolID: user.SchoolID, RegisterTime: user.CreatedAt, Logs: logs})
	}
	util.SetResponse(ctx, rsp)
}
