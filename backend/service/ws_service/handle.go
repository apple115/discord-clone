package ws_service

import (
	"encoding/json"
	"log"
	"time"
)

// ConnectMessage 结构体用于处理连接消息
type ConnectMessage struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

// // JoinChannelMessage 结构体用于处理加入频道消息
// type JoinChannelMessage struct {
// 	ChannelID string `json:"channel_id"`
// }

// SendMessage 结构体用于处理发送消息
type SendMessage struct {
	ChannelID string `json:"channel_id"`
	Content   string `json:"content"`
}

// 连接服务器
func handleConnect(client *Client, data json.RawMessage) {
	var connectMsg ConnectMessage
	err := json.Unmarshal(data, &connectMsg)
	if err != nil {
		log.Printf("消息解析失败: %v", err)
		client.sendError("error")
		return
	}
	accesstoken := connectMsg.Token
	userID := connectMsg.UserID
	ok, err := isValidToken(accesstoken)
	if err != nil {
		client.sendError("error")
		client.Conn.Close()
		return
	}
	if !ok {
		client.sendError("Unauthorized")
		client.Conn.Close()
		return
	}
	// ok, err = isExistUser(userID)
	if ok {
		client.UserID = userID
		client.sendJSON(responceData["connect_success"])
		GuildInstance.addClient(client)
	} else {
		client.sendError("Unauthorized")
		client.Conn.Close()
	}
}

// // 加入频道 更改client的channelID
// func handleJoinChannel(client *Client, data json.RawMessage) {
// 	// 验证client已经连接
// 	if client.UserID == "" {
// 		client.sendError("未授权的连接")
// 		client.Conn.Close()
// 		return
// 	}

// 	var joinChannelMsg JoinChannelMessage
// 	err := json.Unmarshal(data, &joinChannelMsg)
// 	if err != nil {
// 		log.Printf("消息解析失败: %v", err)
// 		client.sendError("消息解析失败")
// 		return
// 	}
// 	if ok, _ := isExistChannel(joinChannelMsg.ChannelID); !ok {
// 		client.sendError("频道不存在")
// 		return
// 	}
// 	//添加 用户client的channelID
// 	client.ChannelID = joinChannelMsg.ChannelID
// 	client.sendJSON(responceData["join_channel_success"])
// }

func handleSendMessage(client *Client, data json.RawMessage) {
	var sendMsg SendMessage
	err := json.Unmarshal(data, &sendMsg)
	if err != nil {
		log.Printf("消息解析失败: %v", err)
		client.sendError("json解析失败")
		return
	}
	if ok, _ := isExistChannel(sendMsg.ChannelID); !ok {
		client.sendError("频道不存在")
		return
	}
	// timestamp := util.GenerateTimestamp()
	Message := map[string]interface{}{
		"time":      time.Now().Unix(),
		"userid":    client.UserID,
		"channelID": sendMsg.ChannelID,
		"message":   sendMsg.Content,
		// "timestamp": timestamp,
	}
	//全服务器广播消息
	GuildInstance.SendJSON(Message)
}

func handleHeartbeat(client *Client, data json.RawMessage) {
	response := map[string]interface{}{
		"type": "heartbeat_ack",
	}
	client.sendJSON(response)
	// 更新客户端的最后心跳时间
	client.lastHeartbeat = time.Now()
}
