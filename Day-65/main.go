package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/", test)

	server.Run(":8888")
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"now": time.Now().Format("2006/01/02 15:04:05")})
}
