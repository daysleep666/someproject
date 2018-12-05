package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//redis 消息队列

func main() {
	for i := 0; i < 3; i++ {
		go func(_index int) {
			for {
				waitMsg(_index)
			}
		}(i)
	}
	var c chan int
	<-c
}

func getRedis() redis.Conn {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return r
}

func waitMsg(_index int) {
	r := getRedis()
	reply, err := r.Do("blpop", "mylist", "0")
	if err != nil {
		panic(err)
	}
	for _, v := range reply.([]interface{}) {
		msg := string(v.([]byte))
		fmt.Println(_index, " says: ", msg)
	}
}
