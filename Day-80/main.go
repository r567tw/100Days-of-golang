package main

import (
    "image"
    "image/color"
    "image/png"
    "os"

    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
)

func main() {
    // 設定圖片的寬度和高度
    width, height := 100, 100

    // 創建一個新的矩形圖片
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // 設置文字顏色和位置
    textColor := color.RGBA {255, 255, 255, 255} // 黑色文字
    text := "HelloWorld"
    point := fixed.Point26_6{
        X: fixed.Int26_6 (1000), // X 軸位置
        Y: fixed.Int26_6 (3000), // Y 軸位置
    }

    // 繪制文字
    d := &font.Drawer{
        Dst:  img,
        Src:  image.NewUniform(textColor),
        Face: basicfont.Face7x13,
        Dot:  point,
    }

    d.DrawString(text)

    // 創建一個檔案來保存圖片
    file, _ := os.Create("result.png")
    defer file.Close()

    // 將圖片以 PNG 格式寫入檔案
    png.Encode(file, img)
}
