package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rusisg/rest-api-gin/handler"
)

func main() {
	r := gin.Default()

	r.GET("/albums", handler.GetAlbums)
	r.Run(":8080")
}
