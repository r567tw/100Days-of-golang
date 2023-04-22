package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int

	fmt.Printf("Please enter the ceiling number for generating a random number: ")
	fmt.Scanf("%d", &num)

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 這一行帶入現在的時間,好讓每一次遊戲隨機產生的數字都不一樣, 原來亂數的原理其實是有一個小技巧和規則的
	result := r.Intn(num)

	fmt.Printf("Random Number: %d", result)
}
