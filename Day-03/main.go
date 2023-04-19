package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://info.cern.ch/"

	res, err := http.Get(url)
	content, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Request is Error")
		fmt.Println(err)
	}

	fmt.Println(res.StatusCode)
	fmt.Println(string(content))
}
