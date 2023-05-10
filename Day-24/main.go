package main

import (
	"fmt"
	"time"
)

func cost(start time.Time) {
	cost := time.Since(start)
	fmt.Printf("time cost = %v\n", cost)
}

func main() {
	defer cost(time.Now())
	temp := 0
	for i := 0; i < 10000000000; i++ {
		temp += 1
	}

	fmt.Println(temp)
}
