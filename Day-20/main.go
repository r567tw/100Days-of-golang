package main

import "fmt"

func main() {
	grades := make(map[string]int, 0)

	grades["jimmy"] = 99
	grades["tom"] = 97
	grades["amy"] = 88

	fmt.Println(grades)
	fmt.Println(len(grades))

	for name, grade := range grades {
		fmt.Printf("%s grade: %d\n", name, grade)
	}

	fmt.Println()

	for name, grade := range grades {
		fmt.Printf("%s grade: %d\n", name, grade)
	}
	// Map is not ordered
	fmt.Println()

	var value int
	var isExist bool

	value, isExist = grades["jimmy"]
	if isExist {
		fmt.Printf("key jimmy is exist, value: %d\n", value)
	} else {
		fmt.Printf("key jimmy not exist\n")
	}

	value, isExist = grades["alice"]
	if isExist {
		fmt.Printf("key alice is exist, value: %d\n", value)
	} else {
		fmt.Printf("key alice not exist\n")
	}
}
