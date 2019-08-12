package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	cmd := client.Get("test1")
	fmt.Println(cmd.Err())
	if cmd.Err().Error() == "redis: nil" {
		fmt.Println("none key")
	}
}
