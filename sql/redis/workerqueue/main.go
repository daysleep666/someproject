package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := getClient()
	go WorkerWatchQueue(client, []Callback{&Buy{Name: "buy"}, &Sell{Name: "sell"}})
	go WorkerDelyQueue(client)
	for i := 0; i < 3; i++ { // 添加3个低优先级的任务，每个任务执行需要三秒钟，一共需要9秒执行完
		Add(client, "low", "buy", 0)
	}
	time.Sleep(time.Second)
	Add(client, "high", "sell", 4) // 添加一个高优先级任务
	time.Sleep(time.Second * 10)
}

// output
// Buy.......
// Sell.......
// Buy.......
// Buy.......

type Task struct {
	Name  string `json:"name"`
	Queue string `json:"queue"`
}

func GetTask(str string) *Task {
	var task Task
	json.Unmarshal([]byte(str), &task)
	return &task
}

func (t *Task) ConvertString() string {
	b, _ := json.Marshal(t)
	return string(b)
}

func Add(r *redis.Client, key, name string, delay int64) {
	if delay == 0 { // 立即执行
		r.RPush(key, name)
	} else {
		task := Task{Name: name, Queue: key}
		r.ZAdd("delayqueue", redis.Z{Member: task.ConvertString(), Score: float64(time.Now().Unix() + delay)})
	}
}

type Callback interface {
	IsMe(string) bool
	Do()
}

type Buy struct {
	Name string
}

func NewBuy() *Buy {
	return &Buy{}
}

func (t *Buy) Do() {
	fmt.Printf("Buy.......\n")
	time.Sleep(time.Second * 3)
}

func (t *Buy) IsMe(name string) bool {
	return t.Name == name
}

type Sell struct {
	Name string
}

func NewSell() *Sell {
	return &Sell{}
}

func (t *Sell) Do() {
	fmt.Printf("Sell.......\n")
}

func (t *Sell) IsMe(name string) bool {
	return t.Name == name
}

func WorkerWatchQueue(r *redis.Client, callBacks []Callback) { // 优先级队列
	for {
		cmd := r.BLPop(time.Second*30, "high", "mid", "low")
		if cmd.Err() != nil {
			continue
		}
		name := cmd.Val()[1]
		for _, v := range callBacks {
			if v.IsMe(name) {
				v.Do()
				break
			}
		}
	}
}

func WorkerDelyQueue(r *redis.Client) {
	for {
		tasks := r.ZRangeByScore("delayqueue", redis.ZRangeBy{Min: "0", Max: fmt.Sprintf("%v", time.Now().Unix()), Count: 100}).Val()
		for _, str := range tasks {
			if task := GetTask(str); task != nil {
				Add(r, task.Queue, task.Name, 0)
				r.ZRem("delayqueue", str)
			}
		}
		if len(tasks) == 100 { // 立刻执行后面的
			continue
		}
		time.Sleep(time.Second)
	}
}

//------------------------------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
