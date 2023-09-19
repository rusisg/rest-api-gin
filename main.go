package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rusisg/rest-api-gin/handler"
)

func main() {
	r := gin.Default()

	r.GET("/albums", handler.GetAlbums)
	r.GET("/albums/:id", handler.GetAlbumByID)
	r.POST("/albums", handler.PostAlbum)
	// r.POST("/albums/", handler.UpdateAlbums) plan to add
	r.DELETE("/albums/:id", handler.DeleteAlbumByID)

	r.Run(":8080")
}
