package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^[a-zA-Z0-9]{4,6}$")

	confirm_true := re.MatchString("abcdef")
	confirm_false := re.MatchString("ab!@df")

	place := regexp.MustCompile("aa")

	replace := place.ReplaceAllString("aabbccaabbccdd", "xx")

	fmt.Println(confirm_true, confirm_false, replace)
}
