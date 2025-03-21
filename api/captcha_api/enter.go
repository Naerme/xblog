package captcha_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct{}

type CaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

var stores = base64Captcha.DefaultMemStore

func (CaptchaApi) CaptchaView(c *gin.Context) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      1,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890",
		//BgColor: &color.RGBA{
		//	R: 3,
		//	G: 102,
		//	B: 214,
		//	A: 125,
		//},//不加默认白色
	}
	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, global.CaptchaStore)
	lid, lb64s, _, err := captcha.Generate()

	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("图片验证码生成失败%s", c)
	}
	res.OkWithData(CaptchaResponse{
		CaptchaID: lid,
		Captcha:   lb64s,
	}, c)
}
