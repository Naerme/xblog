package focus_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type FocusApi struct {
}
type FocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

// FocusUserApi 登录人关注用户
func (FocusApi) FocusUserApi(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)

	claims := jwts.GetClaims(c)
	if cr.FocusUserID == claims.UserID {
		res.FailWithMsg("你时刻都在关注自己", c)
		return
	}
	// 查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("关注用户不存在", c)
		return
	}

	// 查之前是否已经关注过他了
	var focus models.UserFocusModel
	err = global.DB.Take(&focus, "user_id = ? and focus_user_id = ?", claims.UserID, user.ID).Error
	if err == nil {
		res.FailWithMsg("请勿重复关注", c)
		return
	}

	// 每天关注是不是应该有个限度？
	// 每天的取关也要有个限度？

	// 关注
	global.DB.Create(&models.UserFocusModel{
		UserID:      claims.UserID,
		FocusUserID: cr.FocusUserID,
	})

	res.OkWithMsg("关注成功", c)
	return
}
