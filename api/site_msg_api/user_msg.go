package site_msg_api

import (
	"blogx_server/common"
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum/message_type_enum"
	"blogx_server/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserMsgResponse struct {
	CommentMsgCount int `json:"commentMsgCount"`
	DiggMsgCount    int `json:"diggMsgCount"`
	PrivateMsgCount int `json:"privateMsgCount"`
	SystemMsgCount  int `json:"systemMsgCount"`
}

func (SiteMsgApi) UserMsgView(c *gin.Context) {
	claims := jwts.GetClaims(c)
	var msgList []models.MessageModel
	global.DB.Find(&msgList, "rev_user_id = ? and is_read = ?", claims.UserID, false)

	var data UserMsgResponse
	for _, model := range msgList {
		switch model.Type {
		case message_type_enum.CommentType, message_type_enum.ApplyType:
			data.CommentMsgCount++
		case message_type_enum.DiggArticleType, message_type_enum.DiggCommentType, message_type_enum.CollectArticleType:
			data.DiggMsgCount++
		case message_type_enum.SystemType:
			data.SystemMsgCount++
		}
	}

	var chatList []models.ChatModel
	// 接收人是我，而且这个消息未读
	global.DB.Find(&chatList, "rev_user_id = ?", claims.UserID)
	var chatIDList []uint
	for _, model := range chatList {
		chatIDList = append(chatIDList, model.ID)
	}
	chatAcMap := common.ScanMapV2(models.UserChatActionModel{}, common.ScanOption{
		Where: global.DB.Where("chat_id in ?", chatIDList),
	})
	for _, model := range chatList {
		_, ok := chatAcMap[model.ID]
		if !ok {
			data.PrivateMsgCount++
			continue
		}
	}
	// 过滤掉删除的，只取未读的
	var userReadMsgIDList []uint
	global.DB.Model(models.UserGlobalNotificationModel{}).
		Where("user_id = ? and (is_read = ? or is_delete = ?)", claims.UserID, true, true).
		Select("id").Scan(&userReadMsgIDList)
	// 算未读的全局消息
	var systemMsg []models.GlobalNotificationModel
	query := global.DB.Where("")
	if len(userReadMsgIDList) > 0 {
		query.Where("id not in ?", userReadMsgIDList)
	}
	global.DB.Where(query).Find(&systemMsg)
	data.SystemMsgCount += len(systemMsg)

	res.OkWithData(data, c)
}
