package main

import (
	"fmt"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	LIMIT = 1000
)

func main() {

}

func ClearSession(_r *redis.Client) { // 如果token数量超过limit，就删除超过limit的部分
	for {
		count := getTokenCount(_r)
		if count <= LIMIT {
			time.Sleep(time.Second)
			continue
		}
		end := int64(math.Min(float64(count-LIMIT), float64(100)))
		tokens := rangeTokenWithTime(_r, 0, end-1)
		delToken(_r, tokens...)
		delTokenWithTime(_r, tokens...)
		delView(_r, tokens...)
		delCart(_r, tokens...)
	}
}

func UpdateToken(_r *redis.Client, _uID int64, _token string, _itemID int64) {
	setUIDWithToken(_r, _uID, _token) // 维持token-user
	setTokenWithTime(_r, _token)      // 更新令牌最近一次活跃的时间
	if _itemID > 0 {
		setViewWithToken(_r, _token, _itemID)
		remView(_r, _token)
	}
}

// uID-token
func setUIDWithToken(_r *redis.Client, _uID int64, _token string) {
	key := "login:"
	_r.HSet(key, _token, _uID)
}

func getUIDbyToken(_r *redis.Client, _token string) int64 {
	key := "login:"
	cmd := _r.HGet(key, _token)
	n, _ := cmd.Int64()
	return n
}

func delToken(_r *redis.Client, _tokens ...string) {
	key := "login:"
	_r.HDel(key, _tokens...)
}

func setTokenWithTime(_r *redis.Client, _token string) { // 记录token最后一次活跃的时间
	key := "recent"
	_r.ZAdd(key, redis.Z{Member: _token, Score: float64(time.Now().Unix())})
}

func getTokenCount(_r *redis.Client) int64 {
	key := "recent"
	cmd := _r.ZCard(key)
	return cmd.Val()
}

func rangeTokenWithTime(_r *redis.Client, _from, _to int64) []string {
	key := "recent"
	cmd := _r.ZRange(key, _from, _to)
	return cmd.Val()
}

func delTokenWithTime(_r *redis.Client, _tokens ...string) {
	key := "recent"
	_r.ZRem(key, convertToInterfaceArr(_tokens...)...)
}

func setViewWithToken(_r *redis.Client, _token string, _itemID int64) { // 浏览记录
	key := fmt.Sprintf("view:%v", _token)
	_r.ZAdd(key, redis.Z{Member: _itemID, Score: float64(time.Now().Unix())})
}

func remView(_r *redis.Client, _token string) { // 切割浏览记录，只保留前25个
	key := fmt.Sprintf("view:%v", _token)
	_r.ZRemRangeByRank(key, 0, -26)
}

func delView(_r *redis.Client, _tokens ...string) {
	for i, v := range _tokens {
		_tokens[i] = fmt.Sprintf("view:%v", v)
	}
	_r.Del(_tokens...)
}

func addToCart(_r *redis.Client, _token string, _item, _count int64) {
	key := "caty:" + _token
	if _count > 0 {
		_r.HSet(key, fmt.Sprintf("%v", _item), _count)
	} else {
		_r.HDel(key, fmt.Sprintf("%v", _item))
	}
}

func delCart(_r *redis.Client, _tokens ...string) {
	_r.Del(_tokens...)
}

//------------------------------------------------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func convertToInterfaceArr(strs ...string) []interface{} {
	i := make([]interface{}, len(strs))
	for _, v := range strs {
		i = append(i, v)
	}
	return i
}
