package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // 處理 GET 請求
    app.Get("/hello", func(c *fiber.Ctx) error {
        return c.SendString("Hello, GET request!")
    })

    // 處理 POST 請求
    app.Post("/post", func(c *fiber.Ctx) error {
        return c.SendString("POST request received")
    })

    // 處理 PUT 請求
    app.Put("/put", func(c *fiber.Ctx) error {
        return c.SendString("PUT request received")
    })

    // 處理 DELETE 請求
    app.Delete("/delete", func(c *fiber.Ctx) error {
        return c.SendString("DELETE request received")
    })

    app.Listen(":8080")
}
