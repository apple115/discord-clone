package main

import "discord-clone/routers"

func main() {
	r := routers.InitRouter()
	r.Run(":8080")
}
