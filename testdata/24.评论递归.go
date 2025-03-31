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
	// 2 -> 8 -> 9
	rootComment := GetRootComment(3)
	fmt.Println(rootComment.ID)
	rootComment = GetRootComment(2)
	fmt.Println(rootComment.ID)
	rootComment = GetRootComment(1)
	fmt.Println(rootComment.ID)
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
