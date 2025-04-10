// router/site_router.go
package router

import (
	"blogx_server/api"
	"blogx_server/api/data_api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddleware, app.SumView)
	r.GET("data/article", middleware.AdminMiddleware, app.ArticleDataView)
	r.GET("data/register_user", middleware.AdminMiddleware, app.RegisterUserDataView)
	r.GET("data/growth", middleware.AdminMiddleware, middleware.BindQueryMiddleware[data_api.GrowthDataRequest], app.GrowthDataView)

}
