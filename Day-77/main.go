package main

import (
    "fmt"
    "os"
    "image"
    "image/color"
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

    // 創建一個新的圖片，與原始圖片尺寸相同
    editedImg := image.NewRGBA(img.Bounds())

    // 調整亮度
    brightness := 100
    draw.Draw(editedImg, editedImg.Bounds(), img, image.Point{}, draw.Src)
    for y := editedImg.Bounds().Min.Y; y < editedImg.Bounds().Max.Y; y++ {
        for x := editedImg.Bounds().Min.X; x < editedImg.Bounds().Max.X; x++ {
            c := editedImg.At(x, y).(color.RGBA)
            c.R = uint8(clamp(int(c.R) + brightness))
            c.G = uint8(clamp(int(c.G) + brightness))
            c.B = uint8(clamp(int(c.B) + brightness))
            editedImg.Set(x, y, c)
        }
    }

    // 保存修改後的圖片
    output, err := os.Create("edited.png")
    if err != nil {
        fmt.Println ("無法創建輸出文件:", err)
        return
    }
    defer output.Close()

    err = png.Encode(output, editedImg)
    if err != nil {
        fmt.Println ("無法編碼圖片:", err)
        return
    }

    fmt.Println ("圖片處理完成，已保存為 edited.jpg")
}

func clamp(x int) int {
    if x < 0 {
        return 0
    }
    if x > 255 {
        return 255
    }
    return x
}