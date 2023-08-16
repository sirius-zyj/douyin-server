package main

import client "douyin-server/rpc/client"

func main() {
	client.Test()
	// go service.RunMessageServer()
	// dao.Init()
	// r := gin.Default()

	// initRouter(r)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
