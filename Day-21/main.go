package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	answer := rand.Intn(100)

	var number int
	var input string
	fmt.Print("Please input a number (1~100):")
	fmt.Scanf("%s", &input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("please input number")
		os.Exit(0)
	}

	for number != answer {
		if number < answer {
			fmt.Println("too low")
		}

		if number > answer {
			fmt.Println("too high")
		}

		fmt.Print("Please input a number (1~100):")
		fmt.Scanf("%s", &input)
		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("please input number")
			break
		}
	}

	fmt.Printf("Answer is %d\n", answer)

}
