package article_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/redis_service/redis_article"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleCollectRequest struct {
	ArticleID uint `json:"articleID" binding:"required"`
	CollectID uint `json:"collectID"`
}

func (ArticleApi) ArticleCollectView(c *gin.Context) {
	cr := middleware.GetBind[ArticleCollectRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	var collectModel models.CollectModel
	claims := jwts.GetClaims(c)
	fmt.Println("cr.CollectID:", cr.CollectID)
	if cr.CollectID == 0 {
		// 是默认收藏夹
		err = global.DB.Take(&collectModel, "user_id = ? and is_default = ?",
			claims.UserID, 1).Error
		if err != nil {
			// 创建一个默认收藏夹
			collectModel.Title = "默认收藏夹"
			collectModel.UserID = claims.UserID
			collectModel.IsDefault = true
			global.DB.Create(&collectModel)
		}
		cr.CollectID = collectModel.ID
	} else {
		// 判断收藏夹是否存在，并且是否是自己创建的
		err = global.DB.Take(&collectModel, "user_id = ? ", claims.UserID).Error
		if err != nil {
			res.FailWithMsg("收藏夹不存在", c)
			return
		}
	}

	// 判断是否收藏
	var articleCollect models.UserArticleCollectModel
	err = global.DB.Where(models.UserArticleCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		CollectID: cr.CollectID,
	}).Take(&articleCollect).Error

	if err != nil {
		// 收藏
		err = global.DB.Create(&models.UserArticleCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ArticleID,
			CollectID: cr.CollectID,
		}).Error
		if err != nil {
			res.FailWithMsg("收藏失败", c)
			return
		}
		res.OkWithMsg("收藏成功", c)

		// 对收藏夹进行加1
		redis_article.SetCacheCollect(cr.ArticleID, true)
		global.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count + 1"))
		return
	}
	// 取消收藏
	err = global.DB.Where(models.UserArticleCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		CollectID: cr.CollectID,
	}).Delete(&models.UserArticleCollectModel{}).Error

	if err != nil {
		res.FailWithMsg("取消收藏失败", c)
		return
	}
	res.OkWithMsg("取消收藏成功", c)
	redis_article.SetCacheCollect(cr.ArticleID, false)
	global.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count - 1"))
	return
}

func (ArticleApi) ArticleCollectPatchRemoveView(c *gin.Context) {
	var cr = middleware.GetBind[models.RemoveRequest](c)

	claims := jwts.GetClaims(c)

	var userCollectList []models.UserArticleCollectModel

	global.DB.Find(&userCollectList, "id = ? and user_id = ?", cr.IDList, claims.UserID)

	if len(userCollectList) > 0 {
		global.DB.Delete(&userCollectList)
	}
	res.OkWithMsg(fmt.Sprint("批量移除文:章%d", len(userCollectList)), c)

}
