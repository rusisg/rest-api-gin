package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rusisg/rest-api-gin/types"
)

var albums = types.NewAlbums()

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbum(c *gin.Context) {
	var newAlbum types.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
