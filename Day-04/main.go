package main

import (
	"os"
)

func main() {
	f, err := os.Create("./sample.txt")

	if err != nil {
		panic(err)
	}

	f.WriteString("Hello World")
}
