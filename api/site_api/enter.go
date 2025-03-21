package site_api

import (
	"blogx_server/common/res"
	"blogx_server/conf"
	"blogx_server/core"
	"blogx_server/global"
	"blogx_server/middleware"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

type SiteApi struct {
}

type SiteInfoRequset struct {
	Name string `uri:"name"`
}

func (SiteApi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequset
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if cr.Name == "site" {
		global.Conifg.Site.About.Version = global.Version
		res.OkWithData(global.Conifg.Site, c)
	}
	//判断是不是管理员
	middleware.AdminMiddleware(c)
	_, ok := c.Get("claims")
	if !ok {
		return
	}

	//claims, err := jwts.ParseTokenByGin(c)
	//if err != nil {
	//	res.FailWithError(err, c)
	//	return
	//}
	//if claims.Role != enum.AdminRole {
	//	re
	//}

	var data any

	switch cr.Name {
	case "email":
		rep := global.Conifg.Email
		rep.AuthCode = "******"
		data = rep
	case "qq":
		rep := global.Conifg.QQ
		rep.AppKey = "******"
		data = rep
	case "qiNiu":
		rep := global.Conifg.QiNiu
		rep.SecretKey = "******"
		data = rep
	case "ai":
		rep := global.Conifg.Ai
		rep.SecretKey = "******"
		data = rep

	default:
		res.FailWithMsg("不存在的配置1", c)
		return
	}

	res.OkWithData(data, c)

	//log_service.NewLoginSuccess(c, enum.UserPwdLoginType)
	//log_service.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "hjy", "1234")
	//c.JSON(200, gin.H{"code": 0, "msg": "站点信息"})
	return
}

func (SiteApi) SiteInfoQQView(c *gin.Context) {
	res.OkWithData(global.Conifg.QQ.Url(), c)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required" label:"名字"`
	Age  int    `json:"age" binding:"required" label:"年龄"`
}

func (SiteApi) SiteUpdateView(c *gin.Context) {
	//log := log_service.GetLog(c)
	//log.ShowRequest()
	//log.ShowRequestHeader()
	//log.ShowResponseHeader()
	//log.ShowResponse()
	//log.SetTitle("更新站点")
	//log.SetItemInfo("请求时间", time.Now())
	//log.SetImage("/xxx/xxx")
	//log.SetLink("yaml学习地址", "http://www.fengfengzhidao.com")
	//c.Header("xxx", "xxee")
	var cr SiteInfoRequset
	err := c.ShouldBindUri(&cr)
	if err != nil {
		//log.SetError("参数绑定失败", err)
		res.FailWithError(err, c)
		return
	}
	//log.SetItemInfo("结构体", cr)
	//log.SetItemInfo("切片", []string{"a", "b"})
	//log.SetItemInfo("字符串", "你好")
	//log.SetItemInfo("数字", 123)
	//id := log.Save()
	//fmt.Println(1, id)
	//
	//id = log.Save()
	//fmt.Println(2, id)
	//c.JSON(200, gin.H{"code": 0, "msg": "123站点信息", "data": 1})

	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiNiu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	switch s := rep.(type) {
	case conf.Site:
		// 判断站点信息更新前端文件部分
		err = UpdateSite(s)
		if err != nil {
			res.FailWithError(err, c)
			return
		}
		global.Conifg.Site = s
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Conifg.Email.AuthCode
		}
		global.Conifg.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Conifg.QQ.AppKey
		}
		global.Conifg.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Conifg.QiNiu.SecretKey
		}
		global.Conifg.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Conifg.Ai.SecretKey
		}
		global.Conifg.Ai = s
	}

	// 改配置文件
	core.SetConf()

	res.OkWithMsg("更新站点配置成功", c)
	return
}
func UpdateSite(site conf.Site) error {
	if site.Project.Icon == "" && site.Project.Title == "" &&
		site.Seo.Keywords == "" && site.Seo.Description == "" &&
		site.Project.WebPath == "" {
		return nil
	}

	if site.Project.WebPath == "" {
		return errors.New("请配置前端地址")
	}

	file, err := os.Open(site.Project.WebPath)
	if err != nil {
		return errors.New(fmt.Sprintf("%s 文件不存在", site.Project.WebPath))
	}

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		logrus.Errorf("goquery 解析失败 %s", err)
		return errors.New("文件解析失败")
	}

	if site.Project.Title != "" {
		doc.Find("title").SetText(site.Project.Title)
	}
	if site.Project.Icon != "" {
		selection := doc.Find("link[rel=\"icon\"]")
		if selection.Length() > 0 {
			selection.SetAttr("href", site.Project.Icon)
		} else {
			// 没有就创建
			doc.Find("head").AppendHtml(fmt.Sprintf("<link rel=\"icon\" href=\"%s\">", site.Project.Icon))
		}
	}
	if site.Seo.Keywords != "" {
		selection := doc.Find("meta[name=\"keywords\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Keywords)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"keywords\" content=\"%s\">", site.Seo.Keywords))
		}
	}
	if site.Seo.Description != "" {
		selection := doc.Find("meta[name=\"description\"]")
		if selection.Length() > 0 {
			selection.SetAttr("content", site.Seo.Description)
		} else {
			doc.Find("head").AppendHtml(fmt.Sprintf("<meta name=\"description\" content=\"%s\">", site.Seo.Description))
		}
	}

	html, err := doc.Html()
	if err != nil {
		logrus.Errorf("生成html失败 %s", err)
		return errors.New("生成html失败")
	}

	err = os.WriteFile(site.Project.WebPath, []byte(html), 0666)
	if err != nil {
		logrus.Errorf("文件写入失败 %s", err)
		return errors.New("文件写入失败")
	}
	return nil
}
