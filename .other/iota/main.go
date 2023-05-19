package main

import "fmt"

func main() {
	const (
		_ = iota
		TypeA
		TypeB
	)

	fmt.Println(TypeA)
	fmt.Println(TypeB)
}

// Ref: https://www.readfog.com/a/1657316285844918272
