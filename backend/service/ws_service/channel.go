package ws_service

import "sync"

type Channel struct {
	sync.RWMutex
	clients map[string]*Client
}

// removeClient
func (c *Channel) remove(userid string) {
	c.RLock()
	defer c.RUnlock()
	delete(c.clients, userid)
}

// addClient
func (c *Channel) add() {
	c.RLock()
	defer c.RUnlock()
}

func (c *Channel) count() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.clients)
}

func (c *Channel) broadcast(data map[string]interface{}) {
	c.RLock()
	defer c.RUnlock()
	for _, client := range c.clients {
		client.sendJSON(data)
	}
}
