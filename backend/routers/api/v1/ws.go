package v1

import (
	"discord-clone/services"

	"github.com/gin-gonic/gin"
)

func WSHandler(c *gin.Context) {
	services.HandleWebsocket(c)
}
