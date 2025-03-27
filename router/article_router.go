package router

import (
	"blogx_server/api"
	"blogx_server/api/article_api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.ArticleCreateRequest], app.ArticleCreateView)
	r.GET("article", middleware.BindQueryMiddleware[article_api.ArticleListRequest], app.ArticleListView)
}
