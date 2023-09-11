package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rusisg/rest-api-gin/types"
	"golang.org/x/exp/slices"
)

var albums = types.NewAlbums()

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbum(c *gin.Context) {
	var newAlbum types.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for s, a := range albums {
		if a.ID == id {
			albums = slices.Delete(albums, s)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album can't delete"})

}
