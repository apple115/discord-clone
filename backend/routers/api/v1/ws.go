package v1

import (
	"discord-clone/service"

	"github.com/gin-gonic/gin"
)

func WSHandler(c *gin.Context) {
	service.HandleWebsocket(c)
}
