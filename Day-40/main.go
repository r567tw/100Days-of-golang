package main

import "fmt"

func bubble(list []int) []int {
	length := len(list)
	
	// 在Go語言中，slice是引用類型，當你將一個slice作為參數傳遞給函數時，
	// 實際上傳遞的是該slice的引用（指標），而不是複製一份新的slice。這表示在函數內部對slice的修改會影響到原始的slice。
	// Create new variable sorted , 避免直接修改 list 
	sorted := make([]int, length)
	copy(sorted,list)
	// 第幾回合
	index := length
	// times := 1

	for index>0 {
		index --;
		for i:=0; i<index; i++{
			if sorted[i] > sorted[i+1] {
				// 交換
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
			}
		}
		// 過程：
		// fmt.Printf("%d : %v\n", times, sorted)
		// times ++
	}

	return sorted
}


func main(){
	list := []int{3,7,1,9,7,10,8,13,11,12}
	
	result := bubble(list)
	
	fmt.Println("Original: ",list)
	fmt.Println("Result: ",result)
}