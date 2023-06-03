package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8")
	defer db.Close()

	// db.Exec("truncate users")

	tx, err := db.Begin() // 開始一個交易

	if err != nil {
		fmt.Println(err.Error())
	}

	// db.Exec("truncate users")
	db.Exec("INSERT INTO users (id,name) VALUES (?,?)", 1, "jimmy")
	db.Exec("INSERT INTO users (id,name) VALUES (?,?)", 2, "jim")
	db.Exec("INSERT INTO users (id,name) VALUES (?,?)", 3, "jim")
	// goto A
	if err = tx.Commit(); err != nil { // 提交
		fmt.Println(err.Error())
	} else {
		tx.Rollback()
	}

	rows, err := db.Query("select count(*) from users")

	if err != nil {
		fmt.Print(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("count: %d\n", count)
	}

}
