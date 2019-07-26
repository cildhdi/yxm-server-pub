package mp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/config"
	"server/models"
	"server/util"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type (
	loginParam struct {
		Code string `json:"code" binding:"required"`
	}

	loginResponse struct {
		ID    uint   `json:"ID"`
		Token string `json:"Token"`
	}
)

//DirectLogin 直接登录
func DirectLogin(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	util.SetResponse(ctx, &loginResponse{ID: user.ID, Token: user.WxLocalToken})
	return
}

//Login wx登录
func Login(ctx *gin.Context) {
	param := loginParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	response, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + config.AppID + "&secret=" + config.AppSecret + "&js_code=" + param.Code + "&grant_type=authorization_code")
	if err != nil {
		util.SetError(ctx, util.CodeConvertFailed, err.Error())
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.SetError(ctx, util.ReadCodeResponseFailed, err.Error())
		return
	}

	var res interface{}
	json.Unmarshal(body, &res)
	m := res.(map[string]interface{})

	var errcode float64
	openID, sessionKey := "", ""

	for k, v := range m {
		if k == "errcode" {
			if value, ok := v.(float64); ok {
				errcode = value
			}
		} else if k == "openid" {
			if value, ok := v.(string); ok {
				openID = value
			}
		} else if k == "session_key" {
			if value, ok := v.(string); ok {
				sessionKey = value
			}
		}
	}

	if errcode != 0 || len(openID) == 0 || len(sessionKey) == 0 {
		util.SetError(ctx, util.CodeConvertFailed, "code2session fail")
		return
	}

	token, err := uuid.NewV4()
	if err != nil {
		util.SetError(ctx, util.GenerateUUIDFailed, err.Error())
		return
	}

	user := models.User{}
	models.Db().First(&user, "wx_open_id = ?", openID)

	if user.ID == 0 {
		user = models.User{
			WxOpenID:      openID,
			WxSessionKey:  sessionKey,
			WxLocalToken:  token.String(),
			LastLoginDate: time.Now(),
			Name:          "",
			SchoolID:      "",
		}

		if err := models.Db().Create(&user).Error; err != nil {
			util.SetError(ctx, util.CreateUserFailed, err.Error())
			return
		}
	} else {
		if err := models.Db().Model(&user).Updates(models.User{
			WxLocalToken:  token.String(),
			WxSessionKey:  sessionKey,
			LastLoginDate: time.Now(),
		}).Error; err != nil {
			util.SetError(ctx, util.UpdateUserFailed, err.Error())
			return
		}
	}

	util.SetResponse(ctx, &loginResponse{ID: user.ID, Token: token.String()})
	return
}

//UserResponse 用户信息
type UserResponse struct {
	ID       uint   `json:"ID"`
	Name     string `json:"Name"`
	SchoolID string `json:"SchoolID"`
}

//GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	util.SetResponse(ctx, &UserResponse{ID: user.ID, Name: user.Name, SchoolID: user.SchoolID})
}

type setUserInfoParam struct {
	Name     string `json:"Name" binding:"required"`
	SchoolID string `json:"SchoolID" binding:"required"`
}

//SetUserInfo 设置用户信息
func SetUserInfo(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	param := setUserInfoParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	models.Db().Model(&user).Updates(models.User{
		Name:     param.Name,
		SchoolID: param.SchoolID,
	})
	util.SetResponse(ctx, &UserResponse{ID: user.ID, Name: user.Name, SchoolID: user.SchoolID})
}
