package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}

	val, err := client.Get("key").Result()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("key2", val2)
	}
}
