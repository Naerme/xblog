// router/site_router.go
package router

import (
	"blogx_server/api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddleware, app.SumView)
}
