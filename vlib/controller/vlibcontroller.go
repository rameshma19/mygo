package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "ListVideos"})
}

func AddVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Add videos"})
}

func UpdateVideos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Update videos"})
}
