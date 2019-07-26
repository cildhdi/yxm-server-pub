package main

import (
	"server/config"
	"server/middlewares"
	"server/routers/admin"
	"server/routers/mp"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	basicAuth := gin.BasicAuth(gin.Accounts{
		config.AdminUsername: config.AdminPassword,
	})

	adminRouter := router.Group("/admin", basicAuth)
	adminRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	adminRouter.Static("/", "dist")

	api := router.Group("/api")

	apiMp := api.Group("/mp")
	apiMp.POST("/login", mp.Login)
	auth := apiMp.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/direct_login", mp.DirectLogin)

		auth.POST("/userinfo", mp.GetUserInfo)
		auth.POST("/setuserinfo", mp.SetUserInfo)

		auth.POST("/publish", mp.Publish)
		auth.POST("/article", mp.GetArticle)

		auth.POST("/punch", mp.Punch)
		auth.POST("/punches", mp.Punches)
	}

	apiAdmin := api.Group("/admin", basicAuth)
	{
		apiAdmin.POST("/readlog_count", admin.ReadLogCount)
		apiAdmin.POST("/readlogs", admin.ReadLogs)
		apiAdmin.POST("/userinfo", admin.UserInfo)
		apiAdmin.POST("/user_count", admin.UserCount)
		apiAdmin.POST("/users", admin.AllUser)
		apiAdmin.POST("/article_count", admin.ArticleCount)
		apiAdmin.POST("/articles", admin.AllArticle)
		apiAdmin.POST("/publish", mp.Publish)
		apiAdmin.POST("/delete_article", admin.ArticleDelete)
	}
	if gin.Mode() != gin.DebugMode {
		autotls.Run(router, "yxm.cildhdi.cn")
	} else {
		router.Run(":8080")
	}
}
