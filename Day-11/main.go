package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	var u User
	u = User{Name: "Jimmy", Age: 29}

	byteData, err := json.MarshalIndent(u, "", "    ") // or json.Marshal(u)

	if err != nil {
		fmt.Println("json err")
	}

	jsonFile, err := os.Create("./example.json")

	if err != nil {
		fmt.Println("create json file error")
	}

	size, err := jsonFile.Write(byteData)

	if err != nil {
		fmt.Println("write file error")
	}

	defer jsonFile.Close()

	fmt.Printf("file Wrote %d bytes\n", size)
}
