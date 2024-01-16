package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // 處理 GET 請求，讀取查詢參數
    app.Get("/query", func(c *fiber.Ctx) error {
        query := c.Query("param")
        return c.SendString("查詢參數 'param' 的值是: " + query)
    })

    // 處理 POST 請求，讀取 JSON 請求體
    app.Post("/json", func(c *fiber.Ctx) error {
        type Request struct {
            Name string `json:"name"`
        }

        var request Request
        if err := c.BodyParser(&request); err != nil {
            return err
        }

        return c.JSON(fiber.Map{
            "received": request.Name,
        })
    })

    app.Listen(":8080")
}
