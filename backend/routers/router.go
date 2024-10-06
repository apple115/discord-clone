package routers

import (
	v1 "discord-clone/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/ws", v1.WSHandler)
	return r
}
