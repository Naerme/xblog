package chat_api

import (
	"blogx_server/common/res"
	"blogx_server/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (ChatApi) ChatView(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil || claims == nil {
		res.FailWithMsg("请登录", c)
		return
	}

	conn, err := UP.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// 消息类型，消息，错误
		t, p, err := conn.ReadMessage()
		if err != nil {
			// 一般是客户端断开 // websocket: close 1005 (no status)
			fmt.Println(err)
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("你说的是：%s吗？", string(p))))
		fmt.Println(t, string(p))
	}
	defer conn.Close()
	fmt.Println("服务关闭")
}
