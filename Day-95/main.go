package main

import (
    "database/sql"
    "github.com/gofiber/fiber/v2"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func main() {
    app := fiber.New()

    // 開啟 SQLite 資料庫
    database, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    // 創建表格
    statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
    statement.Exec()

    // 新增數據
    app.Post("/user", func(c *fiber.Ctx) error {
        name := c.FormValue("name")
        statement, _ := database.Prepare("INSERT INTO users (name) VALUES (?)")
        statement.Exec(name)
        return c.SendString("User successfully created")
    })

    // 獲取數據
    app.Get("/user/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        row := database.QueryRow("SELECT name FROM users WHERE id = ?", id)
        var name string
        row.Scan(&name)
        return c.SendString("User: " + name)
    })

    app.Listen(":8080")
}
