package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMicro())
	fmt.Println(rand.Intn(100))
}
