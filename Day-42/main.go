package main

import (
	"fmt"
)

// 合併排序主函數
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)
}

// 合併兩個已排序的數組
func merge(left, right []int) []int {
	result := make([]int, 0)
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	result = append(result, left...)
	result = append(result, right...)
	return result
}

func main() {
	arr := []int{9, 5, 1, 8, 3, 7, 4, 6, 2}
	fmt.Println("原始數組:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("合併排序後:", sortedArr)
}
