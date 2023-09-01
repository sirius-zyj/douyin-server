package main

import (
	"douyin-server/controller"
	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

func main() {
	// go service.RunMessageServer()

	client.InitRpcClient()

	r := gin.Default()

	controller.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
