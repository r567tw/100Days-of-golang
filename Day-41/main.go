package main

import "fmt"

func sort(list []int) []int {
	length := len(list)

	sorted := make([]int, length)
	copy(sorted, list)

	// times := 1

	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if sorted[min] > sorted[j] {
				min = j
			}
		}
		if min != i {
			sorted[i], sorted[min] = sorted[min], sorted[i]
		}
		// 過程：
		// fmt.Printf("%d : %v\n", times, sorted)
		// times++
	}

	return sorted
}

func main() {
	list := []int{3, 7, 1, 9, 7, 10, 8, 13, 11, 12}

	result := sort(list)

	fmt.Println("Original: ", list)
	fmt.Println("Result: ", result)
}
