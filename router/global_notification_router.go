package router

import (
	"blogx_server/api"
	"blogx_server/api/global_notification_api"
	"blogx_server/middleware"
	"blogx_server/models"
	"github.com/gin-gonic/gin"
)

func GlobalNotificationRouter(r *gin.RouterGroup) {
	app := api.App.GlobalNotificationApi
	r.POST("global_notification", middleware.AdminMiddleware, middleware.BindJsonMiddleware[global_notification_api.CreateRequest], app.CreateView)
	r.GET("global_notification", middleware.AuthMiddleware, middleware.BindQueryMiddleware[global_notification_api.ListRequest], app.ListView)
	r.DELETE("global_notification", middleware.AdminMiddleware, middleware.BindJsonMiddleware[models.RemoveRequest], app.RemoveAdminView)
	r.POST("global_notification/user", middleware.AuthMiddleware, middleware.BindJsonMiddleware[global_notification_api.UserMsgActionRequest], app.UserMsgActionView)

}
