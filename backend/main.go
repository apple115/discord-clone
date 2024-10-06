package main

import (
	"discord-clone/models"
	"discord-clone/pkg/setting"
	"discord-clone/routers"
)

func main() {
	setting.Setup()
	models.Setup()
	r := routers.InitRouter()
	r.Run(":8080")
}
