package main

import "github.com/gin-gonic/gin"

func main() {
	// CREATING SERVER INSTANCE
	server := gin.Default()
	// RUNNING THE SERVER ON PORT 8080 ON LOCAL HOST
	server.Run(":8080")
}
