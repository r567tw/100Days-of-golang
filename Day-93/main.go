package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
    "time"
)

// 日誌記錄中間件
func loggingMiddleware(c *fiber.Ctx) error {
    start := time.Now()
    log.Println("開始處理請求: ", c.Path())

    err := c.Next() // 繼續執行堆棧中的下一個中間件或路由處理器

    // 請求處理完畢後記錄
    log.Printf("請求處理完成，耗時 %v", time.Since(start))
    return err
}

func main() {
    app := fiber.New()


    // 使用中間件
    app.Use(loggingMiddleware)

    // 定義一個簡單的路由
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    app.Listen(":8080")
}
