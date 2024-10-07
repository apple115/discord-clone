package routers

import (
	v1 "discord-clone/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ws", v1.WSHandler)
	return r
}
