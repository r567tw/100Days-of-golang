package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string

	fmt.Printf("Write your name : ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Hello , %s\n", strings.Title(name))
}
