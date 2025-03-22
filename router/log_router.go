package router

import (
	"blogx_server/api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func LogRouter(rr *gin.RouterGroup) {
	app := api.App.LogApi
	r := rr.Group("").Use(middleware.AdminMiddleware)
	r.GET("logs", app.LogListView)
	r.GET("logs/:id", app.LogReadView)
	r.DELETE("logs", app.LogRemoveView)

	//r.GET("logs", middleware.AdminMiddleware, app.LogListView)
	//r.GET("logs/:id", middleware.AdminMiddleware, app.LogReadView)
	//r.DELETE("logs", middleware.AdminMiddleware, app.LogRemoveView)

}
