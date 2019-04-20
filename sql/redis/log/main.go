package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelError LogLevel = "error"
)

func main() {
	c := getClient()
	for i := 0; i < 1; i++ {
		AddLog(c, LogLevelDebug, "xxapi", "hello")
		AddLog(c, LogLevelDebug, "xxapi", "hi")
		AddLog(c, LogLevelDebug, "xxapi", "youxi")
		AddLog(c, LogLevelDebug, "xxapi", "xixi")
	}
	logs := GetLog(c, LogLevelDebug, "xxapi")
	fmt.Println(logs)
}

func AddLog(c *redis.Client, logLevel LogLevel, name, msg string) {
	key := fmt.Sprintf("[log]%v:%v", name, logLevel)
	log := fmt.Sprintf("[%v]%v", time.Now().String(), msg)
	p := c.Pipeline()
	p.LPush(key, log)
	p.LTrim(key, 0, 99) // 只保留前一百个
	p.Exec()
}

func GetLog(c *redis.Client, logLevel LogLevel, name string) []string {
	key := fmt.Sprintf("[log]%v:%v", name, logLevel)
	cmd := c.LRange(key, 0, -1)
	return cmd.Val()
}

//----------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
