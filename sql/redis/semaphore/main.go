package main

// 计数信号量

import (
	"fmt"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/go-redis/redis"
)

func main() {
	wg := sync.WaitGroup{}
	c := getClient()
	for i := 0; i < 13; i++ {
		wg.Add(1)
		go func(i int) {
			if r := AcquireSemaphore(c, "semname", 10); r != "" {
				fmt.Printf("%v success\n", i)
			} else {
				fmt.Println("failed")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// 会因为两个系统的时间不同引起问题
func AcquireSemaphore(r *redis.Client, semname string, limit int) string {
	identifier := bson.NewObjectId().Hex() //	生成唯一标识符
	p := r.Pipeline()
	p.ZRemRangeByScore(semname, "0", fmt.Sprintf("%v", time.Now().Unix()-10))       // 0 移除过期的信号量
	p.ZAdd(semname, redis.Z{Member: identifier, Score: float64(time.Now().Unix())}) // 1
	p.Exec()

	nn := r.ZRank(semname, identifier)
	if int(nn.Val()) < limit {
		return identifier
	}
	r.ZRem(semname, identifier)
	return ""
}

func ReleaseSemaphore(r *redis.Client, semname string, identifier string) {
	r.ZRem(semname, identifier)
}

// 公平锁
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

func getId(r *redis.Client) int {
	cmd := r.Incr("semnameid")
	return int(cmd.Val())
}

//----------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
