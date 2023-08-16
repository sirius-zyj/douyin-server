package main

import (
	"douyin-server/rpc/client"
)

func main() {
	client.InitRpcClient()
	client.GetVideoByUserId(1)
	// go service.RunMessageServer()
	// dao.Init()
	// r := gin.Default()

	// initRouter(r)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
