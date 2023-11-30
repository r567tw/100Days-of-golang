package main

import "fmt"

func main() {

	var num1, num2 float64
	var result float64
	var operator string

	fmt.Print("請輸入第一個數字： ")
	fmt.Scanln(&num1)

	fmt.Print("請輸入第二個數字： ")
	fmt.Scanln(&num2)

	fmt.Print("請輸入運算符號(+、-、*、/): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		panic("錯誤的運算符號！")
	}

	fmt.Printf("%.0f %s %.0f = %.0f\n", num1, operator, num2, result)
}
