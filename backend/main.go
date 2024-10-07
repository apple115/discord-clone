package main

import (
	"discord-clone/models"
	"discord-clone/pkg/setting"
	"discord-clone/routers"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routers := routers.InitRouter()
	writeTimeout := setting.ServerSetting.WriteTimeout
	readTimeout := setting.ServerSetting.ReadTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           ":8000",
		Handler:        routers,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
