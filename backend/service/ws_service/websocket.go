package ws_service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	GuildInstance = Guild{} // 保存所有连接的客户端，string 为用户 ID
	upgrader      = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	messageTypes = map[string]func(client *Client, data json.RawMessage){
		"connect":      handleConnect,
		// "join_channel": handleJoinChannel,
		"send_message": handleSendMessage,
		"heartbeat":    handleHeartbeat,
	}
)

func HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	UserId, exist := c.Get("user_id")
	if !exist {
		return
	}
	if err != nil {
		log.Println("升级为 WebSocket 失败:", err)
		return
	}
	defer conn.Close()
	// 这里就绑定了用户 ID
	client := &Client{Conn: conn, Send: make(chan []byte), UserID: UserId.(string)}

	go client.writePump() // 发送消息
	client.readPump()     // 读取消息
}
