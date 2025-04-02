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

	r.GET("article/digg/:id", middleware.AuthMiddleware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleDiggView)
	r.POST("article/collect", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.ArticleCollectRequest], app.ArticleCollectView)
	r.DELETE("article/collect", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleCollectPatchRemoveView)

	r.POST("article/examine", middleware.AdminMiddleware, middleware.BindJsonMiddleware[article_api.ArticleExamineRequest], app.ArticleExamineView)

	r.DELETE("article/:id", middleware.AuthMiddleware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleRemoveUserView)
	r.DELETE("article", middleware.AdminMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleRemoveView)

	r.POST("article/history", middleware.BindJsonMiddleware[article_api.ArticleLookRequest], app.ArticleLookView)
	r.GET("article/history", middleware.AuthMiddleware, middleware.BindQueryMiddleware[article_api.ArticleLookListRequest], app.ArticleLookListView)
	r.DELETE("article/history", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleLookRemoveView)

	r.POST("category", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.CategoryCreateRequest], app.CategoryCreateView)
	r.GET("category", middleware.BindQueryMiddleware[article_api.CategoryListRequest], app.CategoryListView)
	r.DELETE("category", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CategoryRemoveView)

	r.POST("collect", middleware.AuthMiddleware, middleware.BindJsonMiddleware[article_api.CollectCreateRequest], app.CollectCreateView)
	r.GET("collect", middleware.BindQueryMiddleware[article_api.CollectListRequest], app.CollectListView)
	r.DELETE("collect", middleware.AuthMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CollectRemoveView)

	r.GET("category/options", middleware.AuthMiddleware, app.CategoryOptionsView)
	r.GET("article/tag/options", middleware.AuthMiddleware, app.ArticleTagOptionsView)
}
