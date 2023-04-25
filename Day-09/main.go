package main

import "fmt"

type g interface {
	int16 | int32 | int64 | float32 | float64 | string
}

func add[T g](a, b T) T {
	return a + b
}

func main() {
	fmt.Printf("%s\n", add("Hello", "World"))
	fmt.Printf("%02d\n", add(int64(3), int64(4)))
	fmt.Printf("%.2f\n", add(3.3, 4.4))
}
