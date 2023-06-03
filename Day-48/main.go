package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := sql.Open("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8")
	defer db.Close()

	_, err := db.Exec("truncate users")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = db.Exec(
		"INSERT INTO users (id,name) VALUES (?,?)",
		1,
		"jimmy",
	)

	if err != nil {
		fmt.Print(err.Error())
	}

	rows, err := db.Query("select id,name from users where name=?", "jimmy")

	if err != nil {
		fmt.Print(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s id is %d\n", name, id)
	}
}
