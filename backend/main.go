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

func InitSetting() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	InitSetting()
	r := routers.InitRouter()
	writeTimeout := setting.ServerSetting.WriteTimeout
	readTimeout := setting.ServerSetting.ReadTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	server := &http.Server{
		Addr:         endPoint,
		Handler:      r,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
