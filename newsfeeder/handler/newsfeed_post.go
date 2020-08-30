package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygo/newsfeeder/datastore"
)

//NewsfeedPostReq function
type NewsfeedPostReq struct {
	Title    string `json:"title" binding:"required"`
	NewsPost string `json:"post" binding:"required"`
}

//NewsfeedPost function
func NewsfeedPost(feed datastore.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody NewsfeedPostReq
		err := c.ShouldBindJSON(&reqBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		item := datastore.Item{
			Title:    reqBody.Title,
			NewsPost: reqBody.NewsPost,
		}
		fmt.Println("printing params: ", item)
		feed.AddItem(item)
		c.Status(http.StatusNoContent)
	}
}
