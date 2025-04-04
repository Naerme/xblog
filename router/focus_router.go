package router

import (
	"blogx_server/api"
	"blogx_server/api/focus_api"
	"blogx_server/middleware"
	"github.com/gin-gonic/gin"
)

func FocusRouter(r *gin.RouterGroup) {
	app := api.App.FocusApi

	r.POST("focus", middleware.AuthMiddleware, middleware.BindJsonMiddleware[focus_api.FocusUserRequest], app.FocusUserView)
	r.GET("focus/my_user", middleware.AuthMiddleware, middleware.BindQueryMiddleware[focus_api.FocusUserListRequest], app.FocusUserListView)

}
