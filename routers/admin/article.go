package admin

import (
	"server/models"
	"server/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//AllArticle 所有文章
func AllArticle(ctx *gin.Context) {
	articles := []models.Article{}
	models.Db().Order("time desc").Find(&articles)
	util.SetResponse(ctx, &articles)
}

//ArticleCount 文章数量
func ArticleCount(ctx *gin.Context) {
	count := 0
	models.Db().Model(&models.Article{}).Count(&count)
	util.SetResponse(ctx, count)
}

//ArticleDelete 删除文章
func ArticleDelete(ctx *gin.Context) {
	param := idParam{}
	if err := ctx.ShouldBindBodyWith(&param, binding.JSON); err != nil {
		util.SetError(ctx, util.ParamError, err.Error())
		return
	}

	article := models.Article{}
	models.Db().First(&article, param.ID)

	if article.ID == 0 {
		util.SetError(ctx, util.ArticleNotFound, "no such article")
		return
	}

	models.Db().Delete(&article)
	util.SetResponse(ctx, &article)
}
