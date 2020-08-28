package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mygo/vlib/controller"
)

func main() {
	server := gin.Default()

	server.GET("/", controller.ListVideos)
	server.Run(":8080")
}
