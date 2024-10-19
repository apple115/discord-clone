package ws_service

import (
	"encoding/json"
	"sync"
)

type Guild struct {
	Users map[string]*Client
	sync.RWMutex
}

// sendJson 向连接的所有在线用户发送json数据
func (g *Guild) SendJSON(data map[string]interface{}) {
	message, _ := json.Marshal(data)
	for _, user := range g.Users {
		user.Send <- message
	}
}

// 添加client
func (g *Guild) addClient(user *Client) {
	g.Lock()
	g.Users[user.UserID] = user
	g.Unlock()
}

// 移除client
func (g *Guild) removeClient(user *Client) {
	g.Lock()
	delete(g.Users, user.UserID)
	g.Unlock()
}
