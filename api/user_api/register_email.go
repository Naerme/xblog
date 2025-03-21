package user_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/jwts"
	"blogx_server/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type RegisterEmailRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) RegisterEmailView(c *gin.Context) {
	var cr RegisterEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	//fmt.Println(cr, cr.EmailID)
	//value, ok := global.EmailVerifyStore.Load(cr.EmailID)
	//
	//fmt.Println(ok)
	//fmt.Println(value)
	//if !ok {
	//	res.FailWithMsg("邮箱验证失败1", c)
	//	return
	//}
	//fmt.Println("2")
	//info, ok := value.(email_store.EmailStoreInfo)
	//if !ok {
	//	res.FailWithMsg("邮箱验证失败2", c)
	//	return
	//}
	//if info.Code != cr.EmailCode {
	//	global.EmailVerifyStore.Delete(cr.EmailID)
	//	res.FailWithMsg("邮箱验证失败3", c)
	//	return
	//}
	//global.EmailVerifyStore.Delete(cr.EmailID)

	if !global.Conifg.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}
	//// 创建用户
	uname := base64Captcha.RandText(5, "0123456789")
	//
	_email, _ := c.Get("email")
	email := _email.(string)
	//
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)

	var user = models.UserModel{
		Username:       fmt.Sprintf("b_%s", uname),
		Nickname:       "邮箱用户",
		RegisterSource: enum.RegisterEmailSourceType,
		Password:       hashPwd,
		Email:          email,
		Role:           enum.UserRole,
	}

	err = global.DB.Create(&user).Error
	if err != nil {
		res.FailWithMsg("邮箱注册失败", c)
		logrus.Errorf("创建用户失败 %s", err)
		return
	}

	// 颁发token
	token, err := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
	if err != nil {
		res.FailWithMsg("邮箱登录失败", c)
		return
	}
	//user_service.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)
}
