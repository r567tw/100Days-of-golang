package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	ID        int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32  `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt string `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func main() {
	//Connection
	conn, err := gorm.Open(mysql.Open("test:test@tcp(localhost:3306)/test?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	// //產生table
	// conn.Debug().AutoMigrate(&Customer{})
	// //判斷有沒有table存在
	// migrator := conn.Migrator()
	// has := migrator.HasTable(&Customer{})
	// //has := migrator.HasTable("GG")
	// if !has {
	// 	fmt.Println("table not exist")
	// }

	//create
	// customer := Customer{Username: "tester", Password: "Password", Status: 1}
	// result := conn.Create(&customer)

	// if result.Error != nil {
	// 	fmt.Println("fail")
	// }

	// Select
	// var customers []Customer
	var tester Customer
	// res := conn.Find(&customers)
	// fmt.Println(res.RowsAffected)
	conn.First(&tester)
	fmt.Println(tester)

	// Update
	// conn.First(&tester)
	// tester.Password = "GG2"
	// conn.Save(&tester)

	// conn.Model(&Customer{}).Where("id = ?", 1).Update("password", "helloGG")
}

// Ref: https://ithelp.ithome.com.tw/articles/10245308
