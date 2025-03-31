package main

import (
	"blogx_server/core"
	"blogx_server/flags"
	"blogx_server/global"
	"blogx_server/models"
	"fmt"
)

func main() {
	flags.Parse()
	global.Conifg = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	//// 2 -> 8 -> 9
	//rootComment := GetRootComment(3)
	//fmt.Println(rootComment.ID)
	//rootComment = GetRootComment(2)
	//fmt.Println(rootComment.ID)
	//rootComment = GetRootComment(1)
	//fmt.Println(rootComment.ID)

	//model := models.CommentModel{
	//	Model: models.Model{ID: 2},
	//}
	//GetCommentTree(&model)

	model := GetCommentTreeV3(1)
	fmt.Println(model.ID)
	for _, c1 := range model.SubCommentList {
		fmt.Println("1:", c1.ID)
		for _, c2 := range c1.SubCommentList {
			fmt.Println("2:", c2.ID)
			for _, c3 := range c2.SubCommentList {
				fmt.Println("3:", c3.ID)
				for _, c4 := range c3.SubCommentList {
					fmt.Println("4:", c4.ID)
				}
			}
		}
	}
}

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

// 查一个评论下的子评论
// 评论树
func GetCommentTree(model *models.CommentModel) {
	global.DB.Preload("SubCommentList").Take(model)
	for _, commentModel := range model.SubCommentList {
		GetCommentTree(commentModel)
	}
}

func GetCommentTreeV3(id uint) (model *models.CommentModel) {
	model = &models.CommentModel{
		Model: models.Model{ID: id},
	}
	global.DB.Preload("SubCommentList").Take(model)
	for i := 0; i < len(model.SubCommentList); i++ {
		commentModel := model.SubCommentList[i]
		item := GetCommentTreeV3(commentModel.ID)
		model.SubCommentList[i] = item
	}
	return
}
