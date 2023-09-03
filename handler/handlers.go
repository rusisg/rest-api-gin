package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rusisg/rest-api-gin/types"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, types.NewAlbum)
}
