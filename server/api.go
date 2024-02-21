package server

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

// CreateApi function is responsible for handling the API routes for our application
func CreateApi(r *gin.Engine) {
	r.GET("/", handlers.LandingHandler)
	r.POST("/api/v1/new", handlers.NewUrlHandler)
	r.GET("/api/v1/:url", handlers.GetUrlHandler)
}
