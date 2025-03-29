package router

import (
	"blogx_server/api"
	"blogx_server/api/article_api"
	"blogx_server/middleware"
	"blogx_server/models"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.ArticleCreateRequest], app.ArticleCreateView)
	r.PUT("article", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.ArticleUpdateRequest], app.ArticleUpdateView)
	r.GET("article", middleware.BindQueryMiddleware[article_api.ArticleListRequest], app.ArticleListView)
	r.GET("article/:id", middleware.BindUriMiddleware[models.IDRequest], app.ArticleDetailView)
	r.POST("article/examine", middleware.AdminMiddleware, middleware.BindJsonMiddleware[article_api.ArticleExamineRequest], app.ArticleExamineView)
}
