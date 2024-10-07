package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn      *websocket.Conn
	UserID    string
	ChannelID string
	Send      chan []byte
}

var clients = make(map[string][]*Client) //全局的连接池，存储每个频道的用户

// 读取消息并处理事件
func (client *Client) readPump() {
	defer func() {
		client.Conn.Close()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			fmt.Println("读取 WebSocket 消息失败:", err)
			break
		}

		// 解析接收到的消息
		var msg map[string]interface{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			fmt.Println("消息解析失败:", err)
			continue
		}

		// 根据消息类型处理不同的事件
		switch msg["type"] {
		case "connect":
			client.handleConnect(msg["data"].(map[string]interface{}))
		case "join_channel":
			client.handleJoinChannel(msg["data"].(map[string]interface{}))
		case "send_message":
			client.handleSendMessage(msg["data"].(map[string]interface{}))
		case "heartbeat":
			client.handleHeartbeat()
		default:
			client.sendError("Unknown event type")
		}
	}
}

// 客户端写入消息到 WebSocket
func (client *Client) writePump() {
	defer func() {
		client.Conn.Close()
	}()
	for {
		message, ok := <-client.Send
		if !ok {
			client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		client.Conn.WriteMessage(websocket.TextMessage, message)
	}
}

// 处理连接事件
func (client *Client) handleConnect(data map[string]interface{}) {
	// 验证用户身份，例如检查JWT token
	token := data["token"].(string)
	userID := data["user_id"].(string)

	if token == "user_jwt_token" { // 简单的身份验证
		client.UserID = userID
		response := map[string]interface{}{
			"type": "connect_ack",
			"data": map[string]interface{}{
				"status":  "success",
				"message": "Connected successfully",
			},
		}
		client.sendJSON(response)
	} else {
		client.sendError("Unauthorized")
		client.Conn.Close()
	}
}

// 处理加入频道事件
func (client *Client) handleJoinChannel(data map[string]interface{}) {
	channelID := data["channel_id"].(string)
	client.ChannelID = channelID

	// 将用户添加到频道的客户端列表
	clients[channelID] = append(clients[channelID], client)

	response := map[string]interface{}{
		"type": "join_channel_ack",
		"data": map[string]interface{}{
			"channel_id": channelID,
			"message":    "Joined channel successfully",
		},
	}
	client.sendJSON(response)
}

// 处理发送消息事件
func (client *Client) handleSendMessage(data map[string]interface{}) {
	message := map[string]interface{}{
		"type": "new_message",
		"data": map[string]interface{}{
			"channel_id": client.ChannelID,
			"user_id":    client.UserID,
			"message_id": "abc123", // 示例消息ID
			"content":    data["content"],
			"timestamp":  time.Now().Format(time.RFC3339),
		},
	}
	messageBytes, _ := json.Marshal(message)

	// 广播消息给同一频道的所有客户端
	for _, c := range clients[client.ChannelID] {
		c.Send <- messageBytes
	}
}

// 处理心跳事件
func (client *Client) handleHeartbeat() {
	response := map[string]interface{}{
		"type": "heartbeat_ack",
	}
	client.sendJSON(response)
}

// 发送错误消息
func (client *Client) sendError(message string) {
	response := map[string]interface{}{
		"type": "error",
		"data": map[string]interface{}{
			"message": message,
		},
	}
	client.sendJSON(response)
}

// 发送 JSON 格式的消息
func (client *Client) sendJSON(data map[string]interface{}) {
	message, _ := json.Marshal(data)
	client.Send <- message
}

func HandleWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("升级为 WebSocket 失败:", err)
		return
	}
	defer conn.Close()
	client := &Client{Conn: conn, Send: make(chan []byte)}

	go client.writePump()
	client.readPump()
}
