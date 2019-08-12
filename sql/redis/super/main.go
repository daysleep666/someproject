package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// put get del ...... range

// redis k8s 自己写一致性哈希, 自动伸缩

// redis k8s压力测试

type KvClient interface {
	// Ping() error
	GetString(key string) (string, error)
	GetBytes(key string) ([]byte, error)
	GetInt(key string) (int, error)
	GetBool(key string) (bool, error)
	GetFloat64(key string) (float64, error)
	SetVal(key string, val interface{}, expiration time.Duration) error
	Exist(keys ...string) (bool, error)
	// TTL(key string) (time.Duration, error)
}

var _ KvClient = (*RedisKvClient)(nil)

func GetRedisClient() KvClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &RedisKvClient{client: client}
}

//redis (redis k8s) tikv

type RedisKvClient struct {
	client *redis.Client
}

func (c *RedisKvClient) Ping() error {
	if c.client == nil {
		return fmt.Errorf("client not init")
	}
	cmd := c.client.Ping()
	if cmd == nil {
		return fmt.Errorf("cmd is nil")
	}
	if cmd.Val() != "PONG" {
		return fmt.Errorf("ping failed")
	}
	return nil
}

func (c *RedisKvClient) GetString(key string) (string, error) {
	if c.client == nil {
		return "", fmt.Errorf("client not init")
	}
	cmd := c.client.Get(key)
	if cmd == nil {
		return "", fmt.Errorf("cmd is nil")
	}
	return cmd.Val(), nil
}

func (c *RedisKvClient) GetBytes(key string) ([]byte, error) {
	if c.client == nil {
		return []byte{}, fmt.Errorf("client not init")
	}
	cmd := c.client.Get(key)
	if cmd == nil {
		return []byte{}, fmt.Errorf("cmd is nil")
	}
	return cmd.Bytes()
}

func (c *RedisKvClient) GetInt(key string) (int, error) {
	if c.client == nil {
		return 0, fmt.Errorf("client not init")
	}
	cmd := c.client.Get(key)
	if cmd == nil {
		return 0, fmt.Errorf("cmd is nil")
	}

	return cmd.Int()
}

func (c *RedisKvClient) GetBool(key string) (bool, error) {
	if c.client == nil {
		return false, fmt.Errorf("client not init")
	}
	cmd := c.client.Get(key)
	if cmd == nil {
		return false, fmt.Errorf("cmd is nil")
	}
	r, err := cmd.Int()
	if err != nil {
		return false, err
	}
	return r == 1, nil
}

func (c *RedisKvClient) GetFloat64(key string) (float64, error) {
	if c.client == nil {
		return 0, fmt.Errorf("client not init")
	}
	cmd := c.client.Get(key)
	if cmd == nil {
		return 0, fmt.Errorf("cmd is nil")
	}
	return cmd.Float64()
}
func (c *RedisKvClient) SetVal(key string, value interface{}, expiration time.Duration) error {
	if c.client == nil {
		return fmt.Errorf("client not init")
	}
	cmd := c.client.Set(key, value, expiration)
	return cmd.Err()
}
func (c *RedisKvClient) Exist(keys ...string) (bool, error) {
	if c.client == nil {
		return false, fmt.Errorf("client not init")
	}
	cmd := c.client.Exists(keys...)
	if cmd == nil {
		return false, fmt.Errorf("cmd is nil")
	}
	return cmd.Val() == 1, nil
}
func (c *RedisKvClient) TTL(key string) (time.Duration, error) {
	if c.client == nil {
		return 0, fmt.Errorf("client not init")
	}
	cmd := c.client.TTL(key)
	if cmd == nil {
		return 0, fmt.Errorf("cmd is nil")
	}
	return cmd.Val(), nil
}

func main() {
	client := GetRedisClient()
	fmt.Println(client.Ping())
	err := client.SetVal("hello", "world", 0)
	if err != nil {
		panic(err)
	}
	val, err := client.GetString("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
