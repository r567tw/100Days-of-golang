package main

import (
    "fmt"
    "os"
    "image"
    _ "image/png"
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

    fmt.Println ("圖片寬度:", img.Bounds ().Dx ())
    fmt.Println ("圖片高度:", img.Bounds ().Dy ())
}