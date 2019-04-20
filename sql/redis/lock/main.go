package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var num int

func DoSomething(c *redis.Client, lockName string) {
	var locked int
	if locked = acquireLock(c, lockName); locked == -1 { // 获取锁失败了
		// fmt.Printf("oh shit\n")
		return
	}
	// fmt.Printf("i'm coming %v n = %v\n", locked, num)
	if num > 0 {
		fmt.Printf("第 %v 抢到了\n", num)
		num--
	}
	// fmt.Printf("i'm going %v\n", locked)
	releaseLock(c, lockName, locked)
}

func acquireLock(c *redis.Client, lockName string) int {
	iden := rand.Intn(1000000)
	end := time.Now().Unix() + 10 // 最多尝试十秒获得锁
	for time.Now().Unix() < end {
		cmd := c.SetNX(lockName, iden, 5*time.Second)
		if cmd.Val() {
			return iden
		}
		time.Sleep(time.Nanosecond * 100)
	}
	return -1
}

func releaseLock(c *redis.Client, lockName string, iden int) {
	err := c.Watch(func(tx *redis.Tx) error { // 确保这把锁除了我没别人动过
		if c.Get(lockName).Val() == fmt.Sprintf("%v", iden) { // 确保这是我的锁
			t := tx.TxPipeline()
			t.Del(lockName)
			t.Exec()
		}
		return nil
	}, lockName)
	if err == nil {
		return
	}
}

func main() {
	num = 10
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		c := getClient()
		go func() {
			DoSomething(c, "add")
			wg.Done()
		}()
		// go func() { num++ }()
	}
	wg.Wait()
}

//----------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
