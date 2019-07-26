package mp

import (
	"server/models"
	"server/util"
	"time"

	"github.com/gin-gonic/gin"
)

//Punch 打卡
func Punch(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	now := time.Now()
	begin := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	log := models.ReadLog{}
	models.Db().Where("user_id = ? AND time BETWEEN ? AND ?", user.ID, begin, begin.Add(time.Minute*(24*60-1)+time.Second*59)).First(&log)
	if log.ID == 0 {
		log := models.ReadLog{
			UserID: user.ID,
			Time:   time.Now(),
		}
		if err := models.Db().Create(&log).Error; err != nil {
			util.SetError(ctx, util.CreateReadLogFailed, err.Error())
			return
		}
		util.SetResponse(ctx, &log)
	} else {
		util.SetError(ctx, util.AlreadyPunched, "already punched")
	}
}

//Punches 打卡记录
func Punches(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	logs := []models.ReadLog{}
	models.Db().Where("user_id = ?", user.ID).Find(&logs)

	util.SetResponse(ctx, &logs)
}
