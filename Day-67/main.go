package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"net/http"

	"day-67/models"
)

var DB *gorm.DB

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

	DB = conn

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	r.GET("/customers", getCustomers)


	r.Run(":8888")
}

func getCustomers(context *gin.Context) {

	var customers []models.Customer
	DB.Find(&customers)

	context.JSON(http.StatusOK, gin.H{"data": customers})
}
// https://medium.com/skyler-record/go-%E5%AF%A6%E4%BD%9C-restful-api-2a32210adeaf
// https://ithelp.ithome.com.tw/articles/10245889
// https://blog.logrocket.com/rest-api-golang-gin-gorm/
