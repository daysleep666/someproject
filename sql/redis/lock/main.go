package main

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-redis/redis"
)

var num int

func DoSomething(c *redis.Client, lockName string) {
	var locked string
	if locked = acquireLock(c, lockName); locked == "" { // 获取锁失败了
		// fmt.Printf("oh shit\n")
		return
	}
	// fmt.Printf("i'm coming %v n = %v\n", locked, num)
	if num > 0 {
		fmt.Printf("第 %v 抢到了\n", num)
		num--
		// time.Sleep(time.Second * 3)
	}
	// fmt.Printf("i'm going %v\n", locked)
	releaseLock(c, lockName, locked)
}

func acquireLock(c *redis.Client, lockName string) string {
	iden := bson.NewObjectId().Hex()
	end := time.Now().Unix() + 10 // 最多尝试十秒获得锁
	for time.Now().Unix() < end {
		cmd := c.SetNX(lockName, iden, 5*time.Second)
		if cmd.Val() {
			return iden
		}
		time.Sleep(time.Nanosecond * 100)
	}
	return ""
}

func releaseLock(c *redis.Client, lockName string, iden string) {
	err := c.Watch(func(tx *redis.Tx) error { // 确保这把锁除了我没别人动过
		if c.Get(lockName).Val() == iden { // 确保这是我的锁
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
	for i := 0; i < 20; i++ {
		c := getClient()
		wg.Add(1)
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
