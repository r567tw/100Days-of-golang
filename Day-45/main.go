package main

import "fmt"

func search(list []int, key int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		middle := left + (right-left)/2

		if key > list[middle] {
			left = middle + 1
		}
		if key < list[middle] {
			right = middle - 1
		}
		if key == list[middle] {
			return middle
		}
	}

	return -1

}

func main() {
	var arr = []int{1, 3, 5, 7, 9, 12, 13, 16, 19, 23, 26}

	keyword := 12
	index := search(arr, keyword)

	if index != -1 {
		fmt.Printf("在%d位置上找到%d\n", index, keyword)
	} else {
		fmt.Printf("找不到%d\n", keyword)
	}

	keyword = 14
	index = search(arr, keyword)

	if index != -1 {
		fmt.Printf("在%d位置上找到%d\n", index, keyword)
	} else {
		fmt.Printf("找不到%d\n", keyword)
	}
}
