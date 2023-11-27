package main

import (
    "fmt"
    "os"
    "image"
    "image/draw"
    "image/png"
)

func main() {
    file, err := os.Open("ex.png")
    if err != nil {
        fmt.Println ("無法打開圖片:", err)
        return
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println ("無法解碼圖片:", err)
        return
    }

    wmb, _ := os.Open("watermark.png")
    watermark, _ := png.Decode(wmb)
    defer wmb.Close()


    // 加入浮水印    
    offset := image.Pt (512, 512) // 位置
    b := img.Bounds () // 邊長寬度
    newImage := image.NewRGBA (b) // 創建一個新的圖片，與原始圖片尺寸相同
    draw.Draw(newImage, b, img, image.ZP, draw.Src)
    draw.Draw(newImage, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)


    // 保存修改後的圖片
    output, _ := os.Create("result.png")
    defer output.Close()

    png.Encode(output, newImage)
    fmt.Println ("圖片處理完成，已保存為 result.png")
}