package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := getClient()
	cmd := client.Get("hello")

	fmt.Println(cmd.Val())
}

func Vote(_r *redis.Client, _id string, _userID int) {
	if checkArticleHasVoted(_r, _id, _userID) { // 检查是否投票过了
		return
	}

	if ct := int64(getArticleTime(_r, _id)); time.Now().Unix()-ct > 3600 || ct == 0 { // 投票是否截止 || 文章是否存在
		return
	}

	haddArticleScore(_r, _id, 1)         // 增加自己的分数
	addArticleByScore(_r, _id, 1)        // 增加票数
	addArticleHasVoted(_r, _id, _userID) // 增加已投票记录
}

func PostArticle(_r *redis.Client, _id string, _title string) {

}

func getArticleID(_r *redis.Client) int64 {
	cmd := _r.Incr("articleid")
	return cmd.Val()
}

func addArticle(_r *redis.Client, _id string, _title string, _time int) {
	key := "article:" + _id
	_r.HSet(key, "title", _title)
	_r.HSet(key, "time", _time)
}

func haddArticleScore(_r *redis.Client, _id string, _score int64) {
	key := "article:" + _id
	_r.HIncrBy(key, "score", _score)
}

func addArticleByTime(_r *redis.Client, _id string, _time int) {
	key := "articlebytime"
	_r.ZAdd(key, redis.Z{Member: _id, Score: float64(_time)})
}

func getArticleTime(_r *redis.Client, _id string) int {
	key := "articlebytime"
	cmd := _r.ZScore(key, _id)
	return int(cmd.Val())
}

func addArticleByScore(_r *redis.Client, _id string, _score int) {
	key := "articlebyscore"
	_r.ZIncrBy(key, float64(_score), _id)
}

func addArticleHasVoted(_r *redis.Client, _id string, _userID int) bool {
	key := "voted:" + _id
	cmd := _r.SAdd(key, _userID)
	return cmd.Val() == 1
}

func checkArticleHasVoted(_r *redis.Client, _id string, _userID int) bool {
	key := "voted:" + _id
	cmd := _r.SIsMember(key, _userID)
	return cmd.Val()
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
