package chat_api

import (
	"blogx_server/common/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/ctype/chat_msg"
	"blogx_server/models/enum/chat_msg_type"
	"blogx_server/utils/jwts"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var OnlineMap = map[uint]map[string]*websocket.Conn{}

type ChatRequest struct {
	RevUserID uint                  `json:"revUserID"` // 发给谁
	MsgType   chat_msg_type.MsgType `json:"msgType"`   // 1 文本 2 图片  3 md
	Msg       chat_msg.ChatMsg      `json:"msg"`       // 消息主体
}
type ChatResponse struct {
	ChatListResponse
}

func (ChatApi) ChatView(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil || claims == nil {
		res.FailWithMsg("请登录", c)
		return
	}

	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Errorf("ws升级失败 %s", err)
		return
	}

	userID := claims.UserID
	addr := conn.RemoteAddr().String()
	addrMap, ok := OnlineMap[userID]
	if !ok {
		OnlineMap[userID] = map[string]*websocket.Conn{
			addr: conn,
		}
	} else {
		_, ok1 := addrMap[addr]
		if !ok1 {
			OnlineMap[userID][addr] = conn
		}
	}
	fmt.Println("进入", OnlineMap)
	for {
		// 消息类型，消息，错误
		_, p, err1 := conn.ReadMessage()
		if err1 != nil {
			// 一般是客户端断开 // websocket: close 1005 (no status)
			fmt.Println(err)
			break
		}

		var req ChatRequest
		err2 := json.Unmarshal(p, &req)

		if err2 != nil {
			res.SendConnFailWithMsg("参数错误", conn)
			continue
		}
		// 判断接收人在不在
		var revUser models.UserModel
		err = global.DB.Take(&revUser, req.RevUserID).Error
		if err != nil {
			res.SendConnFailWithMsg("接收人不存在", conn)

			continue
		}

		// 先落库
		item := ChatResponse{
			ChatListResponse: ChatListResponse{
				ChatModel: models.ChatModel{
					MsgType: req.MsgType,
					Msg:     req.Msg,
				},
			},
		}
		res.SendWsMsg(OnlineMap, req.RevUserID, item)
		res.SendConnOkWithData(item, conn)

	}
	defer conn.Close()

	addrMap2, ok2 := OnlineMap[userID]
	if ok2 {
		_, ok3 := addrMap2[addr]
		if ok3 {
			delete(OnlineMap[userID], addr)
		}
		if len(addrMap2) == 0 {
			delete(OnlineMap, userID)
		}
	}

	fmt.Println("离开", OnlineMap)
}
