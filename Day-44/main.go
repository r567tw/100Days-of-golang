package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	target := 6
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Target %d at Index %d \n", target, index)
	} else {
		fmt.Printf("Target %d Not Found。\n", target)
	}
}
