package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	type student struct {
		Name     string
		Chinese  int
		English  int
		Computer int
	}

	file, _ := os.Open("./example.csv") //openfile 比較囉唆一點 要給什麼權限設定之類的
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ','

	r.Read() //ignore first row (column)
	record, _ := r.Read()

	chinese, _ := strconv.Atoi(record[1])
	english, _ := strconv.Atoi(record[2])
	computer, _ := strconv.Atoi(record[3])

	var s = student{
		Name:     record[0],
		Chinese:  chinese,
		English:  english,
		Computer: computer,
	}

	fmt.Printf("Name %s \nChinese Score: %d \nEnglish Score: %d \nComputer Score: %d \n",
		s.Name, s.Chinese, s.English, s.Computer)
}
