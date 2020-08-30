package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mygo/newsfeeder/handler"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.PingGet)
	r.Run(":9090") // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
