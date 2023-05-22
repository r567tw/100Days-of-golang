package main

import "fmt"

func main() {
	hashTable := make(map[string]int)

	hashTable["apple"] = 1
	hashTable["banana"] = 2
	hashTable["orange"] = 3

	fmt.Println("apple:", hashTable["apple"])
	fmt.Println("banana:", hashTable["banana"])
	fmt.Println("orange:", hashTable["orange"])

	delete(hashTable, "banana")

	_, exists := hashTable["banana"]
	fmt.Println("banana exists:", exists)
}
