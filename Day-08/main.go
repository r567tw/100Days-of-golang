package main

import "fmt"

func main() {
	s := "Reverse"

	var result []rune

	s_array := []rune(s)

	for i := len(s_array) - 1; i >= 0; i-- {
		result = append(result, s_array[i])
	}

	fmt.Println(string(result))
}
