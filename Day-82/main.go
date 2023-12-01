package main

import (
	"os"
	"fmt"
)


func main() {
	os.Setenv("NEW_VAR", "12345")
	fmt.Println(os.Getenv("NEW_VAR"))
}
