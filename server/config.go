package server

import (
	"github.com/gin-gonic/gin"
)

// ConnectServer starts the server
func ConnectServer() *gin.Engine {
	r := gin.Default()

	r.Use(setHeaders())

	CreateApi(r)

	return r
}

// SetHeaders function sets the required headers for our API handlers
func setHeaders() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json:charset=utf-8")
		ctx.Header("Host", ctx.Request.Host)
		ctx.Header("X-Powered-By", "Golang")
		ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Token, Accept, X-Requested-With")
	}
}
