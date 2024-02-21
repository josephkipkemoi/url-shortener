package main

import (
	_ "url-shortener/database"
	"url-shortener/server"
)

// Application function entry point
func main() {
	r := server.ConnectServer() // start sever
	r.Run(":8000")
}
