package middlewares

import (
	"server/models"
	"server/util"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type userTokenParam struct {
	WxLocalToken string `json:"token" binding:"required"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := userTokenParam{}
		if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
			util.SetError(ctx, util.ParamError, err.Error())
			return
		}

		var user models.User
		models.Db().First(&user, "wx_local_token = ?", param.WxLocalToken)

		if user.ID == 0 {
			util.SetError(ctx, util.UserNotFound, "no such user")
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
