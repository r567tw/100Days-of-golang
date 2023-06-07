package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("connect error", err.Error())
	} else {
		fmt.Println(pong)
	}
}
