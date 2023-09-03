package main

import (
	"context"
	"douyin-server/config"
	"douyin-server/controller"
	"douyin-server/middleware/otel"
	"douyin-server/rpc/client"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	otel.Init(context.Background(), config.RouterOtelName)
	defer otel.Close()

	client.InitRpcClient()

	r := gin.Default()

	controller.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
