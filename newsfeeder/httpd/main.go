package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mygo/newsfeeder/datastore"
	"github.com/mygo/newsfeeder/handler"
)

//This is main function
func main() {
	feed := datastore.New()

	r := gin.Default()
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
