package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"net/http"

	"Day66/models"
)

func main() {
	//Connection
	conn, err := gorm.Open(mysql.Open("test:test@tcp(localhost:3306)/test?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	//判斷有沒有table存在
	migrator := conn.Migrator()
	has := migrator.HasTable(&models.Customer{})
	if !has {
		//產生table
		conn.Debug().AutoMigrate(&models.Customer{})
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	r.Run()
}
