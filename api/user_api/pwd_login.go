package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	var cr PwdLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Conifg.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用密码登录", c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "(username = ? or email = ?)and password <> ''",
		cr.Val, cr.Val).Error
	if err != nil {
		res.FailWithMsg("用户名或者email未找到", c)

	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("密码错误", c)
		return
	}
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Role:     user.Role,
		Username: user.Username,
	})
	res.OkWithData(token, c)
}
