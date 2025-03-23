package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models/enum"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdateUserPasswordRequest struct {
	OldPwd string `json:"oldPwd" binding:"required"`
	Pwd    string `json:"pwd" binding:"required"`
}

func (UserApi) UpdatePasswordView(c *gin.Context) {
	var cr UpdateUserPasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	claims := jwts.GetClaims(c)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if !(user.RegisterSource == enum.RegisterEmailSourceType || user.Email != "") {
		res.FailWithMsg("不支持修改密码，仅支持邮箱注册或绑定邮箱的修改密码", c)
		return
	}
	//校验密码
	if !pwd.CompareHashAndPassword(user.Password, cr.OldPwd) {
		res.FailWithMsg("旧密码错误", c)
		return
	}
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.OkWithMsg("密码更新成功", c)
	return
}
