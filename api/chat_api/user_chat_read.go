package chat_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ChatApi) ChatReadView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	var chat models.ChatModel
	err := global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMsg("消息不存在", c)
		return
	}

	claims := jwts.GetClaims(c)
	var chatAc models.UserChatActionModel
	err = global.DB.Take(&chatAc, "user_id = ? and chat_id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		global.DB.Create(&models.UserChatActionModel{
			UserID: claims.UserID,
			ChatID: cr.ID,
			IsRead: true,
		})
		res.OkWithMsg("消息读取成功", c)
		return
	}

	if chatAc.IsDelete {
		res.FailWithMsg("消息被删除", c)
		return
	}
	fmt.Println("已读取")
	global.DB.Model(&chatAc).Update("is_read", true)

	res.OkWithMsg("消息读取成功", c)
}
