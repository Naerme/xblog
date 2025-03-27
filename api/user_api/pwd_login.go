package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/service/user_service"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PwdLoginRequest struct {
	Val      string `json:"val" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	cr := middleware.GetBind[PwdLoginRequest](c)
	if !global.Conifg.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用密码登录", c)
		return
	}
	var user models.UserModel
	fmt.Println(user)
	err := global.DB.Take(&user, "(username = ? or email = ?)and password <> ''",
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
	//fmt.Println(token)
	//fmt.Println(user.Role)
	user_service.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)
}
