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
