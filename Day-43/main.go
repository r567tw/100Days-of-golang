package main

import (
	"fmt"
)

func quickSort(data []int) {
	if len(data) <= 1 {
		return
	}

	pivot := data[0]
	left := []int{}
	right := []int{}

	for _, num := range data[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(data, append(append(left, pivot), right...))
}

func main() {
	arr := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	fmt.Println("原始數組:", arr)

	quickSort(arr)
	fmt.Println("排序後:", arr)
}
