package article_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/redis_service/redis_article"

	//"blogx_server/service/redis_service/redis_article"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleLookRequest struct {
	ArticleID  uint `json:"articleID" binding:"required"`
	TimeSecond int  `json:"timeSecond"` // 读文章一共用了多久
}

func (ArticleApi) ArticleLookView(c *gin.Context) {
	cr := middleware.GetBind[ArticleLookRequest](c)

	// TODO: 未登录用户，浏览量如何加
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.OkWithMsg("未登录", c)
		return
	}
	//
	// 引入缓存
	// 当天这个用户请求这个文章之后，将用户id和文章id作为key存入缓存，在这里进行判断，如果存在就直接返回
	//if redis_article.GetUserArticleHistoryCache(cr.ArticleID, claims.UserID) {
	//	logrus.Infof("在缓存里面")
	//	res.OkWithMsg("成功", c)
	//	return
	//}

	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 查这个文章今天有没有在足迹里面
	var history models.UserArticleLookHistoryModel
	err = global.DB.Take(&history,
		"user_id = ? and article_id = ? and created_at < ? and created_at > ?",
		claims.UserID, cr.ArticleID,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02")+" 00:00:00",
	).Error
	if err == nil {
		res.OkWithMsg("成功", c)
		return
	}

	err = global.DB.Create(&models.UserArticleLookHistoryModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
	}).Error
	if err != nil {
		res.FailWithMsg("失败", c)
		return
	}

	redis_article.SetCacheLook(cr.ArticleID, true)
	//redis_article.SetUserArticleHistoryCache(cr.ArticleID, claims.UserID)

	res.OkWithMsg("成功", c)
}
