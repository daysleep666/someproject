package main

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		client := getClient()
		pipe := client.Pipeline()
		wg.Add(1)
		go func() {
			for j := 0; j < 10000; j++ {
				ic := pipe.Incr("num")
				if ic.Err() != nil {
					fmt.Println(ic.Err())
				}
			}
			pipe.Exec()
			wg.Done()
		}()
	}
	wg.Wait()
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
