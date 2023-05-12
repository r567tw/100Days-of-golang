package main

import (
	"fmt"
	"time"
)

func cost(start time.Time) {
	cost := time.Since(start)
	fmt.Printf("time cost = %v\n", cost)
}

func main() {
	number := 100
	fmt.Println(number)
	now_a := time.Now()
	primes := make(chan bool)

	for i := 2; i <= number; i++ {
		go isPrime(i, primes)
		prime := <-primes
		if prime {
			fmt.Printf("%d ", i)
		}
	}
	cost_a := time.Since(now_a)
	fmt.Println()
	fmt.Printf("cost = %v\n", cost_a)
}

func isPrime(n int, ch chan bool) bool {
	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			ch <- false
			return false
		}
	}
	ch <- true
	return true
}
