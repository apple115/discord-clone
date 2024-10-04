package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}
