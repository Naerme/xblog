package comment_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/service/comment_service"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type CommentCreateRequest struct {
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"articleID" binding:"required"`
	ParentID  *uint  `json:"parentID"`
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	cr := middleware.GetBind[CommentCreateRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "id = ? and status = ?", cr.ArticleID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	claims := jwts.GetClaims(c)

	model := models.CommentModel{
		Content:   cr.Content,
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		ParentID:  cr.ParentID,
	}
	//去找这个评论的根评论
	if cr.ParentID != nil {
		// 有父评论
		r := comment_service.GetRootComment(*cr.ParentID)
		if r != nil {
			model.RootParentID = &r.ID
		}
	}

	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithMsg("发布评论失败", c)
		return
	}

	res.OkWithMsg("发布评论成功", c)

}
