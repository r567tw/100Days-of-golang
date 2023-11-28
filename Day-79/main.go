package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

func main() {
    // 創建一個新的矩形圖片
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))

    // 填充白色背景
    whiteColor := color.RGBA{255, 255, 255, 255}
    draw.Draw(img, img.Bounds(), &image.Uniform{whiteColor}, image.ZP, draw.Src)

    // 創建一個檔案來保存圖片
    file, _ := os.Create("background.png")
    defer file.Close()

    // 將圖片以 PNG 格式寫入檔案
    png.Encode(file, img)
}
