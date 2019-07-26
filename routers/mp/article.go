package mp

import (
	"server/models"
	"server/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type publishParam struct {
	Time    int64  `json:"Time" binding:"required"`
	Content string `json:"Content" binding:"required"`
}

//Publish 发布文章
func Publish(ctx *gin.Context) {
	param := publishParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	article := models.Article{Time: time.Unix(param.Time, 0), Content: param.Content}
	if err := models.Db().Create(&article).Error; err != nil {
		util.SetError(ctx, util.ArticleCreateFailed, err.Error())
		return
	}
	util.SetResponse(ctx, &article)
	return
}

type getArticleParam struct {
	Time int64 `json:"Time" binding:"required"`
}

//GetArticle 获取文章
func GetArticle(ctx *gin.Context) {
	param := getArticleParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	article := models.Article{}
	paramTime := time.Unix(param.Time, 0)
	begin := time.Date(paramTime.Year(), paramTime.Month(), paramTime.Day(), 0, 0, 0, 0, time.Local)
	models.Db().First(&article, "time BETWEEN ? AND ?", begin, begin.Add(time.Minute*(24*60-1)+time.Second*59))

	if article.ID == 0 {
		util.SetError(ctx, util.ArticleNotFound, "article not found")
		return
	}

	util.SetResponse(ctx, &article)
	return
}
