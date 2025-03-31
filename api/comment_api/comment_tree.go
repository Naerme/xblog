package comment_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/comment_service"
	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentTreeView(c *gin.Context) {
	var cr = middleware.GetBind[models.IDRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 把根评论查出来
	var commentList []models.CommentModel
	global.DB.Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]comment_service.CommentResponse, 0)
	for _, model := range commentList {
		response := comment_service.GetCommentTreeV4(model.ID)
		list = append(list, *response)
	}

	res.OkWithList(list, len(list), c)
}
