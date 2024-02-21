package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Golang URL shortener",
	})
}
