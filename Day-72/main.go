package main

import (
	"fmt"
	"time"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// Logger 中間件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 請求前
		t := time.Now()

		// 請求處理
		c.Next()

		// 請求後
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("%d %s in %v\n", status, c.Request.RequestURI, latency)
	}
}

func main() {
	
	r := gin.Default()

	// 全局中間件
	r.Use(Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})
	
	r.Run(":8888")
}

