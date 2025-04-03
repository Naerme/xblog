package flags

import (
	"blogx_server/global"
	"blogx_server/models"
	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
		&models.ArticleModel{},
		&models.ArticleDiggModel{},
		&models.CategoryModel{},
		&models.CollectModel{},
		&models.UserArticleCollectModel{},
		&models.ImageModel{},
		&models.UserTopArticleModel{},
		&models.UserArticleLookHistoryModel{},
		&models.CommentModel{},
		&models.BannerModel{},
		&models.LogModel{},
		&models.UserLoginModel{},
		&models.GlobalNotificationModel{},
		&models.ImageModel{},
		&models.UserLoginModel{},              // 用户登陆记录表
		&models.UserTopArticleModel{},         // 用户置顶文章表
		&models.CommentDiggModel{},            // 用户点赞评论表
		&models.MessageModel{},                // 站内信表
		&models.UserMessageConfModel{},        // 用户消息配置表
		&models.UserGlobalNotificationModel{}, // 用户全局消息表

	)
	if err != nil {
		logrus.Errorf("数据库迁移失败%s", err)
		return
	}
	logrus.Infof("迁移成功")
}
