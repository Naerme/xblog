package chat_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/middleware"
	"blogx_server/models"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (ChatApi) UserChatDeleteView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	var chat models.ChatModel
	err := global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMsg("消息不存在", c)
		return
	}
	claims := jwts.GetClaims(c)
	// 我之前是不是已经删过了
	var chatAc models.UserChatActionModel
	err = global.DB.Take(&chatAc, "user_id = ? and chat_id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		global.DB.Create(&models.UserChatActionModel{
			UserID:   claims.UserID,
			ChatID:   cr.ID,
			IsDelete: true,
		})
		res.OkWithMsg("消息删除成功", c)
		return
	}
	if chatAc.IsDelete {
		// 说明之前删过了
		res.FailWithMsg("消息已被删除", c)
		return
	}

	global.DB.Model(&chatAc).Update("is_delete", true)
	res.OkWithMsg("消息删除成功", c)
}
