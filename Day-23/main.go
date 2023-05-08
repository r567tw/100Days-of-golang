package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2) // O(2^n)
}

func main() {
	var n int
	fmt.Printf("Input your fibonacci n: ")
	fmt.Scanf("%d", &n)

	for i := 1; i < n+1; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Print("\n")
}
