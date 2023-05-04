package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := "60-30"

	var result int
	var operator string

	regex := regexp.MustCompile(`[\+\-\*/]`)
	operator = regex.FindString(input)

	numbers_str := strings.Split(input, operator)

	//covert []string to []int
	var numbers []int
	for _, number_str := range numbers_str {
		number, err := strconv.Atoi(number_str)
		numbers = append(numbers, number)
		if err != nil {
			fmt.Println("error")
			break
		}
	}

	switch operator {
	case "+":
		result = numbers[0] + numbers[1]
	case "-":
		result = numbers[0] - numbers[1]
	case "*":
		result = numbers[0] * numbers[1]
	case "/":
		if numbers[1] == 0 {
			panic("不能被0除")
		}
		result = numbers[0] / numbers[1]
	default:
		fmt.Println("Error")
	}

	fmt.Printf("%s = %d\n", input, result)
	// expect:3+4 = 7
}
