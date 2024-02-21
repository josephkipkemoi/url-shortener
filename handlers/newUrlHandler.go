package handlers

import (
	"net/http"
	"url-shortener/database/models"

	"github.com/gin-gonic/gin"
)

func NewUrlHandler(c *gin.Context) {
	url := models.Url{}
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	url.InsertUrl() // save shortened url and map it to the original url

	c.JSON(http.StatusCreated, gin.H{
		"original_url":  url.Original_Url,
		"shortened_url": url.Shortened_Url,
	})
}
