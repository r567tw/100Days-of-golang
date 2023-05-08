package main

import "fmt"

func sum_of_squares(number int) int {
	if number == 1 {
		return 1
	} else {
		return sum_of_squares(number-1) + number*number
	}
}

func main() {
	var number int
	result := 0

	fmt.Print("Please input a number: ")
	fmt.Scanf("%d", &number)

	if number == 0 {
		fmt.Println("Can't input zero")
		fmt.Println("Please input a number:")
		fmt.Scanf("%d", &number)
	}

	for i := 1; i <= number; i++ {
		result += i * i
	}
	// alse = (number*(number+1))*(2*number+1)/6 a.k.a n*(n+1)*(2n+1)/6
	// Ref: http://www.mathland.idv.tw/fun/Sum_of_squares.htm

	fmt.Printf("The sum of squares from 1 to %d is %d\n", number, result)
	fmt.Printf("Use Recursion: %d\n", sum_of_squares(number))
	fmt.Printf("Use Formula: %d\n", (number*(number+1))*(2*number+1)/6)
}
