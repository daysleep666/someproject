package main

import (
	"time"

	"github.com/go-redis/redis"
)

func main() {
	c := getClient()
	for i := 0; i < 100; i++ {
		AddCount(c, "test")
		time.Sleep(time.Second)
	}
}

func AddCount(c *redis.Client, api string) {
	ts := time.Now().Unix()
	tmp := ts % 5
	ts -= tmp
	if tmp > 5 {
		ts += 5
	}
	c.ZIncr("count:5:hits", redis.Z{Member: ts, Score: 1})
}

//----------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
