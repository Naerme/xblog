package cron_service

import (
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/service/redis_service/redis_article"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SyncArticle() {
	collectMap := redis_article.GetAllCacheCollect()
	diggMap := redis_article.GetAllCacheDigg()
	lookMap := redis_article.GetAllCacheLook()
	commentMap := redis_article.GetAllCacheComment()

	var list []models.ArticleModel
	global.DB.Find(&list)

	for _, model := range list {
		collect := collectMap[model.ID]
		digg := diggMap[model.ID]
		look := lookMap[model.ID]
		comment := commentMap[model.ID]
		if collect == 0 || digg == 0 || look == 0 || comment == 0 {
			continue
		}

		err := global.DB.Model(&model).Updates(map[string]any{
			"look_count":    gorm.Expr("look_count + ?", look),
			"digg_count":    gorm.Expr("digg_count + ?", digg),
			"collect_count": gorm.Expr("collect_count + ?", collect),
			"comment_count": gorm.Expr("comment_count + ?", comment),
		}).Error
		if err != nil {
			logrus.Errorf("更新失败 %s", err)
			continue
		}
		logrus.Infof("%s 更新成功", model.Title)
	}

	// 这里可能会有增量的数据
	// 可以再获取一次
	//collectMap = redis_article.GetAllCacheCollect()
	//diggMap = redis_article.GetAllCacheDigg()
	//lookMap = redis_article.GetAllCacheLook()

	// 走完之后清空掉
	redis_article.Clear()

	// 再同步回去

}
