package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Customer struct {
	ID        int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;" json:"id,omitempty"`
	Username  string `gorm:"type:varchar(20) NOT NULL;" json:"username,omitempty"`
	Password  string `gorm:"type:varchar(100) NOT NULL;" json:"password,omitempty"`
	Status    int32  `gorm:"type:int(5);" json:"status,omitempty"`
	CreatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

func main() {
	//Connection
	conn, err := gorm.Open(mysql.Open("test:test@tcp(localhost:3306)/test?charset=utf8"), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	conn.Exec("TRUNCATE customers")


	customer := Customer{Username: "jimmy", Password: "Password", Status: 1}
	conn.Create(&customer)

	customer2 := Customer{Username: "jimmy2", Password: "Password", Status: 1}
	conn.Create(&customer2)

	conn.Where("id = ?", "1").Delete(&Customer{})

	var count int64
	conn.Model(&Customer{}).Count(&count)
	fmt.Println(count)
}
