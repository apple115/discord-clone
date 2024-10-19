package v1

import (
	"discord-clone/service/ws_service"
	"github.com/gin-gonic/gin"
)

func WSHandler(c *gin.Context) {
	ws_service.HandleWebsocket(c)
}
