package comment_service

import (
	"blogx_server/global"
	"blogx_server/models"
)

// GetRootComment 获取一个评论的根评论
func GetRootComment(commentID uint) (model *models.CommentModel) {
	var comment models.CommentModel
	err := global.DB.Take(&comment, commentID).Error
	if err != nil {
		return nil
	}
	if comment.ParentID == nil {
		// 没有父评论了，那么他就是根评论
		return &comment
	}
	return GetRootComment(*comment.ParentID)
}

// GetCommentTree 获取评论树
func GetCommentTree(model *models.CommentModel) {
	global.DB.Preload("SubCommentList").Take(model)
	for _, commentModel := range model.SubCommentList {
		GetCommentTree(commentModel)
	}
}

// GetCommentTreeV2 获取评论树
func GetCommentTreeV2(id uint) (model *models.CommentModel) {
	model = &models.CommentModel{
		Model: models.Model{ID: id},
	}

	global.DB.Preload("SubCommentList").Take(model)
	for i := 0; i < len(model.SubCommentList); i++ {
		commentModel := model.SubCommentList[i]
		item := GetCommentTreeV2(commentModel.ID)
		model.SubCommentList[i] = item
	}
	return
}
