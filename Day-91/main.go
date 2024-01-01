package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // 創建一個新的 Fiber 實例
    app := fiber.New()

    // 定義一個 GET 路由
    app.Get("/", func(c *fiber.Ctx) error {
        // 返回字符串響應
        return c.SendString("Hello, Fiber!")
    })

    // 啟動 Fiber 應用
    app.Listen(":8080")
}
