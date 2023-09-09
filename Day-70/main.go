package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"net/http"

	"day-70/models"
)

var DB *gorm.DB

func main() {
	//Connection
	conn, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test?charset=utf8"), &gorm.Config{})
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
	r.POST("/customers", createCustomer)
	r.GET("/customers/:id", getCustomer)
	r.PUT("/customers/:id", updateCustomer)


	r.Run(":8888")
}

func getCustomers(context *gin.Context) {
	var customers []models.Customer
	var customerData []struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		Bio  string `json:"bio"`
	}

	DB.Find(&customers)

	// 将查询结果映射到自定义结构
	for _, customer := range customers {
		customerData = append(customerData, struct {
			ID uint `json:"id"`
			Name string `json:"name"`
			Bio  string `json:"bio"`
		}{
			ID: customer.ID,
			Name: customer.Name,
			Bio:  customer.Bio,
		})
	}

	context.JSON(http.StatusOK, gin.H{"data": customerData})
}


func createCustomer(context *gin.Context) {

	type CreateCustomerInput struct {
		Name  string `json:"name" binding:"required"`
		Bio   string `json:"bio" binding:"required"`
	}

	var input CreateCustomerInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create customer
	customer := models.Customer{Name: input.Name,Bio: input.Bio}

	DB.Create(&customer)

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

func getCustomer(context *gin.Context) {
	customerID := context.Param("id")

	var customer models.Customer

	result := DB.First(&customer, customerID)
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

func updateCustomer(context *gin.Context) {
	customerID := context.Param("id")

	var customer models.Customer

	result := DB.First(&customer, customerID)
	
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}


	type UpdateCustomerInput struct {
		Name  string `json:"name" binding:"required"`
		Bio   string `json:"bio" binding:"required"`
	}

	var input UpdateCustomerInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Model(&customer).Updates(input)
	context.JSON(http.StatusOK, gin.H{"data": customer})
}


// ref: https://blog.logrocket.com/rest-api-golang-gin-gorm/