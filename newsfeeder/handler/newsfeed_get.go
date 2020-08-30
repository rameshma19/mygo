package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mygo/newsfeeder/datastore"
)

//NewsfeedGet function
func NewsfeedGet(feed datastore.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAllItems()
		c.JSON(http.StatusOK, results)
	}
}
