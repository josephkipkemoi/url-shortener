package handlers

import (
	"net/http"
	"url-shortener/database/models"

	"github.com/gin-gonic/gin"
)

func GetUrlHandler(c *gin.Context) {
	url, ok := c.Params.Get("url")
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
	}

	s := models.GetUrl(url)

	c.JSON(http.StatusTemporaryRedirect, gin.H{
		"original_url": s,
	})
}
