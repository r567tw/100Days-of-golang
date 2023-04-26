package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData, err := os.Open("./example.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonData.Close()

	byteData, _ := ioutil.ReadAll(jsonData)

	var u User
	json.Unmarshal(byteData, &u)
	fmt.Printf("User: %s, and Age: %d\n", u.Name, u.Age)
}
