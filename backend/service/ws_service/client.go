package ws_service

import (
	"encoding/json"
	"log"
	"time"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn          *websocket.Conn
	UserID        string
	Send          chan []byte // 发送消息的通道
	lastHeartbeat time.Time   // 最后心跳时间
}

// 发送 JSON 格式的消息
func (client *Client) sendJSON(data map[string]interface{}) {
	message, _ := json.Marshal(data)
	client.Send <- message
}

func (client *Client) sendError(message string) {
	client.sendJSON(map[string]interface{}{
		"message": message,
	})
}

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
		messageType := string(msg["type"])
		data := msg["data"]
		if handler, ok := messageTypes[messageType]; ok {
			handler(client, data)
		} else {
			client.sendError("Unknown event type")
		}
		// 检查心跳时间，如果超过一定时间没有心跳，关闭连接
		if time.Since(client.lastHeartbeat) > 100*time.Second {
			log.Println("心跳超时，关闭连接:client:", client.UserID)
			client.Conn.Close()
			break
		}
	}
}

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
