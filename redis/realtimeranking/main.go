package main

import (
	"fmt"
	"math/rand"

	"github.com/garyburd/redigo/redis"
)

// 搞一个redis集合， 测单实例在海量数据下的性能

type myRedis interface {
	redis.Conn
}

func main() {
	r := getRedis()
	for i := 0; i < 100000000; i++ {
		writeMsg(r, i)
	}
}

func getRedis() myRedis {
	r, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	return r
}

func writeMsg(_conn redis.Conn, _index int) {
	key := fmt.Sprintf("user:%v", _index)
	value := fmt.Sprintf("%v", rand.Int63n(10000))
	reply, err := _conn.Do("set", key, value)
	if err != nil {
		panic(err)
	}
	fmt.Println(_index, ":", reply.(string))

}
