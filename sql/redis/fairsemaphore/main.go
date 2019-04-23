package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	wg := sync.WaitGroup{}
	c := getClient()
	for i := 0; i < 13; i++ {
		wg.Add(1)
		go func(i int) {
			identifier := acquireLock(c, "semnamelock")
			if identifier != -1 {
				if r := AcquireFairSemaphore(c, "semname", 10); r != "" {
					fmt.Printf("%v success\n", i)
				} else {
					fmt.Println("failed")
				}
			}
			releaseLock(c, "semnamelock", identifier)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 公平信号量
func AcquireFairSemaphore(r *redis.Client, semname string, limit int) string {
	identifier := bson.NewObjectId().Hex() //	生成唯一标识符
	czet := semname + ":owner"
	ctr := semname + ":counter" // 累加计数器
	p := r.Pipeline()
	p.ZRemRangeByScore(semname, "0", fmt.Sprintf("%v", time.Now().Unix()-10))  // 移除过期的信号量
	p.ZInterStore(czet, redis.ZStore{Weights: []float64{1, 0}}, czet, semname) // 交集 更新下czet中的记录 权重1*czet，0*semname
	p.Exec()
	counter := r.Incr(ctr).Val()
	p.ZAdd(semname, redis.Z{Member: identifier, Score: float64(time.Now().Unix())}) // 用来记录这个identifier是否过期了
	p.ZAdd(czet, redis.Z{Member: identifier, Score: float64(counter)})              // 用来记录这个identifier的累加数
	p.Exec()
	rank := int(r.ZRank(czet, identifier).Val())
	if rank < limit {
		return identifier
	}

	// 获取公平信号量失败
	p.ZRem(semname, identifier)
	p.ZRem(czet, identifier)
	p.Exec()
	return ""
}

func RefreshFairSemaphore(r *redis.Client, semname, identifier string) error { // 更新信号量 防止信号量过期
	cmd := r.ZAddXX(semname, redis.Z{Member: identifier, Score: float64(time.Now().Unix())}) // 需要xx来保证这个信号量没有被干掉
	return cmd.Err()
}

func ReleaseFairSemaphore(r *redis.Client, semname, identifier string) {
	p := r.Pipeline()
	czet := semname + ":owner"
	p.ZRem(semname, identifier)
	p.ZRem(czet, identifier)
	p.Exec()
}

//----------------------------------------------------------------------
// 锁的代码

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

//----------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
