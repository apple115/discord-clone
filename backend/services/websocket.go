package services

import (
	"context"
	"discord-clone/models"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


// ConnectMessage 结构体用于处理连接消息
type ConnectMessage struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// JoinChannelMessage 结构体用于处理加入频道消息
type JoinChannelMessage struct {
	ChannelID string `json:"channel_id"`
}

// SendMessage 结构体用于处理发送消息
type SendMessage struct {
	Content string `json:"content"`
}

// HeartbeatMessage 结构体用于处理心跳消息
type HeartbeatMessage struct{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn          *websocket.Conn
	UserID        string
	ChannelID     string
	Send          chan []byte
	lastHeartbeat time.Time // 最后心跳时间
}

var (
	clients = make(map[string][]*Client) //全局的连接池，存储每个频道的用户
	mu      sync.RWMutex
)

// 读取消息并处理事件
func (client *Client) readPump() {
	defer func() {
		client.Conn.Close()
	}()
	for {
		_, message, err := client.Conn.ReadMessage()
		if err != nil {
			log.Printf("读取 WebSocket 消息失败: %v", err)
			break
		}
		// 解析接收到的消息
		var msg map[string]json.RawMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Printf("消息解析失败: %v", err)
			continue
		}

		// 根据消息类型处理不同的事件
		switch string(msg["type"]) {
		case `"connect"`:
			var connectMsg ConnectMessage
			err = json.Unmarshal(msg["data"], &connectMsg)
			if err != nil {
				log.Printf("消息解析失败: %v", err)
				continue
			}
			client.handleConnect(connectMsg)
		case `"join_channel"`:
			var joinChannelMsg JoinChannelMessage
			err = json.Unmarshal(msg["data"], &joinChannelMsg)
			if err != nil {
				log.Printf("消息解析失败: %v", err)
				continue
			}
			client.handleJoinChannel(joinChannelMsg)
		case `"send_message"`:
			var sendMsg SendMessage
			err = json.Unmarshal(msg["data"], &sendMsg)
			if err != nil {
				log.Printf("消息解析失败: %v", err)
				continue
			}
			client.handleSendMessage(sendMsg)
		case `"heartbeat"`:
			client.handleHeartbeat()
		default:
			client.sendError("Unknown event type")
		}
		// 检查心跳时间，如果超过一定时间没有心跳，关闭连接
		if time.Since(client.lastHeartbeat) > 100*time.Second {
			client.Conn.Close()
			break
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
		// 检查连接是否仍然打开
		if client.Conn.WriteMessage(websocket.PingMessage, nil) != nil {
			log.Printf("连接已关闭，无法发送消息")
			return
		}
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("发送 WebSocket 消息失败: %v", err)
			return
		}
	}
}

// 处理连接事件
// TODO 验证用户身份
func (client *Client) handleConnect(data ConnectMessage) {
	// 验证用户身份，例如检查JWT token
	token := data.Token
	userID := data.UserID

	if isValidToken(token) { // 增强身份验证
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

func isValidToken(token string) bool {
	// 这里可以添加更复杂的身份验证逻辑
	return token == "user_jwt_token"
}


// 处理加入频道事件
// TODO 确认channelID是否存在
func (client *Client) handleJoinChannel(data JoinChannelMessage) {
	channelID := data.ChannelID
	client.ChannelID = channelID

	mu.Lock()
	defer mu.Unlock()

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
func (client *Client) handleSendMessage(data SendMessage) {
	content := data.Content
	message := map[string]interface{}{
		"type": "new_message",
		"data": map[string]interface{}{
			"channel_id": client.ChannelID,
			"user_id":    client.UserID,
			"message_id": "abc123", // 示例消息ID
			"content":    content,
			"timestamp":  time.Now().Format(time.RFC3339),
		},
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("消息序列化失败: %v", err)
		return
	}

	mu.RLock()
	defer mu.RUnlock()

	// 广播消息给同一频道的所有客户端
	for _, c := range clients[client.ChannelID] {
		select {
		case c.Send <- messageBytes:
		default:
			log.Printf("无法发送消息到客户端 %s, 关闭连接", c.UserID)
			close(c.Send)
			removeClient(client.ChannelID, c)
		}
	}

	// 存储消息到 MongoDB
	collection := models.MongoDB.Collection("messages")
	_, err = collection.InsertOne(context.TODO(), message)
	if err != nil {
		log.Println("存储消息到 MongoDB 失败:", err)
	}
}

// 处理心跳事件
func (client *Client) handleHeartbeat() {
	response := map[string]interface{}{
		"type": "heartbeat_ack",
	}
	client.sendJSON(response)

	// 更新客户端的最后心跳时间
	client.lastHeartbeat = time.Now()
}

func removeClient(channelID string, client *Client) {
	mu.Lock()
	defer mu.Unlock()
	for i, c := range clients[channelID] {
		if c == client {
			clients[channelID] = append(clients[channelID][:i], clients[channelID][i+1:]...)
			break
		}
	}
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
