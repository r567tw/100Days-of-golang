package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("../Day-04/sample.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
