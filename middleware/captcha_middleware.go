package middleware

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
)

type CaptchaMiddleWareRequest struct {
	CaptchaID   string `json:"captchaID" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func CaptchaMiddleware(c *gin.Context) {
	//if !global.Conifg.Site.Login.Captcha {
	//	return
	//}
	//body, err := c.GetRawData()
	//if err != nil {
	//	res.FailWithMsg("获取请求体错误", c)
	//	c.Abort()
	//	return
	//}
	//c.Request.Body = io.NopCloser(bytes.NewReader(body)) //body用完写回去
	//var cr CaptchaMiddleWareRequest
	//err = c.ShouldBindJSON(&cr)
	//if err != nil {
	//	res.FailWithMsg("图形验证失败", c)
	//	c.Abort()
	//	return
	//}
	//if !global.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
	//	res.FailWithMsg("验证码错误", c)
	//	c.Abort()
	//	return
	//}
	//c.Request.Body = io.NopCloser(bytes.NewReader(body)) //body用完写回去
	if !global.Conifg.Site.Login.Captcha {
		fmt.Println("已关闭验证码")
		return
	}
	body, err := c.GetRawData()
	if err != nil {
		res.FailWithMsg("获取请求体错误", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	var cr CaptchaMiddleWareRequest
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf("图形验证失败 %s", err)
		res.FailWithMsg("图形验证失败", c)
		c.Abort()
		return
	}
	if !global.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
		res.FailWithMsg("验证码错误", c)
		c.Abort()
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
}
