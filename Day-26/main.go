package main

import (
	"fmt"
	"time"
)

func main() {
	go printA()
	go printB()
	// 有時A先、有時B先, 看作業系統怎麼調用！
	time.Sleep(time.Second * 5)
	fmt.Println("End")
}

func printA() {
	for i := 0; i < 3; i++ {
		fmt.Println("A")
	}
}

func printB() {
	for i := 0; i < 3; i++ {
		fmt.Println("B")
	}
}

// Ref: https://ithelp.ithome.com.tw/articles/10256509
//
// 一般來說使用多執行緒中，最常會遇到會5個問題如下:

// 多執行緒相互溝通
// 等待一執行緒結束後再接續工作
// 多執行緒共用同一個變數
// 不同執行緒產出影響後續邏輯
// 兄弟執行緒間不求同生只求同死
