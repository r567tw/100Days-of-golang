package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	drivers := sql.Drivers()
	fmt.Println(drivers)

	db, _ := sql.Open("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8")

	err := db.Ping()

	if err != nil {
		fmt.Println("Connected False", err.Error())
	} else {
		fmt.Println("Connected Success")
		db.Close()
	}

	db.Close()
}
