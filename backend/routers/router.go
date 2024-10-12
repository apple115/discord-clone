package routers

import (
	"discord-clone/routers/api"
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
	r.POST("/register", api.Register)
	r.POST("/login", api.GetAuth)
	r.POST("/access-token", api.RefreshToken) // 获得新的access token
	r.GET("/login/github", api.GitHubLogin)
	r.GET("/callback", api.GitHubCallback)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use()
	{
		apiv1.GET("/ws", v1.WSHandler)
		// 获取所有频道
		apiv1.GET("/channels", v1.GetChannels)
		// 创建频道
		apiv1.POST("/channels", v1.CreateChannel)

		// 删除频道
		apiv1.DELETE("/channels/:channelID", v1.DeleteChannel)

		// 更新频道信息
		apiv1.PUT("/channels/:channelID", v1.UpdateChannel)

		// 获取特定频道的详细信息
		apiv1.GET("/channels/:channelID", v1.GetChannelByID)

		// 获取特定频道的所有消息
		apiv1.GET("/channels/:channelID/messages", v1.GetChannelMessageByID)

	}
	return r
}
