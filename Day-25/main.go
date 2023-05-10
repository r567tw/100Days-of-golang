package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Print("請輸入三角形想要的階層： ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		spaces := strings.Repeat(" ", n-i)
		stars := strings.Repeat("*", 2*i-1)
		fmt.Printf("%s%s", spaces, stars)
		fmt.Println()
	}
}
