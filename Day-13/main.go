package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("./example.csv")

	writer := csv.NewWriter(f)
	err := writer.WriteAll([][]string{
		{"Name", "Score"},
		{"Jimmy", "100"},
	})

	if err != nil {
		fmt.Println("Write CSV Error !")
	} else {
		fmt.Println("Write CSV Success !")
	}
}
