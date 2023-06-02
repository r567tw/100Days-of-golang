package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	up := true

	m, err := migrate.New(
		"file://migrations",
		"mysql://test:test@tcp(localhost:3306)/test",
	)

	if err != nil {
		panic(err)
	}

	if up {
		err = m.Up()
	} else {
		err = m.Down()
	}

	if err != nil {
		panic(err)
	}
}
