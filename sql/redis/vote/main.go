package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const (
	continuedTime = 3600
	onePage       = 15
)

func main() {
	client := getClient()

	// for i := 0; i < 1000; i++ {
	// 	// id := fmt.Sprintf("%v", i)
	// 	// PostArticle(client, fmt.Sprintf("title:%v", id))
	// 	rand.New(rand.NewSource(time.Now().Unix()))
	// 	Vote(client, fmt.Sprintf("%v", rand.Int63n(20)), rand.Int63n(1000000))
	// }

	// for i := 0; i < 20; i++ {
	// 	rand.New(rand.NewSource(time.Now().Unix()))
	// 	addGroups(client, fmt.Sprintf("%v", rand.Int63n(20)), []string{"1", "2", "3"})
	// }

	getGroupArticle(client, 1, 0)
	// GetArticleTopByScore(client, 0)
}

func Vote(_r *redis.Client, _id string, _userID int64) {
	if checkArticleHasVoted(_r, _id, _userID) { // 检查是否投票过了
		return
	}

	if ct := int64(getArticleTime(_r, _id)); time.Now().Unix()-ct > continuedTime || ct == 0 { // 投票是否截止 || 文章是否存在
		return
	}

	haddArticleScore(_r, _id, 1)         // 增加自己的分数
	addArticleByScore(_r, _id, 1)        // 增加票数
	addArticleHasVoted(_r, _id, _userID) // 增加已投票记录
}

func PostArticle(_r *redis.Client, _title string) string {
	id := getArticleID(_r)
	expireArticleHasVoted(_r, id)                 //  投票持续一周
	addArticle(_r, id, _title, time.Now().Unix()) // 增加文章
	addArticleByTime(_r, id, time.Now().Unix())   // 增加文章 时间排序
	return id
}

func GetArticleTopByScore(_r *redis.Client, _page int64) {
	ids := getArticleIDsTop(_r, "articlebyscore", _page)
	GetArticleByIDs(_r, ids)
}

func GetArticleByIDs(_r *redis.Client, _ids []string) {
	for i, id := range _ids {
		key := "article:" + id
		a := getArticleByKey(_r, key)
		fmt.Printf("第%v篇文章\n", i)
		fmt.Printf("	title:%v\n", a["title"])
		fmt.Printf("	score:%v\n", a["score"])
		fmt.Printf("	time:%v\n", a["time"])
		fmt.Println()
	}

}

// 文章
func getArticleID(_r *redis.Client) string {
	cmd := _r.Incr("articleid")
	return fmt.Sprintf("%v", cmd.Val())
}

func addArticle(_r *redis.Client, _id string, _title string, _time int64) {
	key := "article:" + _id
	_r.HSet(key, "title", _title)
	_r.HSet(key, "time", _time)
}

func getArticleByKey(_r *redis.Client, _key string) map[string]string {
	cmd := _r.HGetAll(_key)
	return cmd.Val()
}

func haddArticleScore(_r *redis.Client, _id string, _score int64) {
	key := "article:" + _id
	_r.HIncrBy(key, "score", _score)
}

// 时间排行榜
func addArticleByTime(_r *redis.Client, _id string, _time int64) {
	key := "articlebytime"
	_r.ZAdd(key, redis.Z{Member: _id, Score: float64(_time)})
}

func getArticleTime(_r *redis.Client, _id string) int64 {
	key := "articlebytime"
	cmd := _r.ZScore(key, _id)
	return int64(cmd.Val())
}

func getArticleIDsTop(_r *redis.Client, _key string, _page int64) []string {
	start := _page * onePage
	end := start + onePage
	cmd := _r.ZRevRange(_key, start, end)
	return cmd.Val()
}

// 分数排行榜
func addArticleByScore(_r *redis.Client, _id string, _score int64) {
	key := "articlebyscore"
	_r.ZIncrBy(key, float64(_score), _id)
}

func expireArticleHasVoted(_r *redis.Client, _id string) {
	key := "voted:" + _id
	_r.Expire(key, time.Second*continuedTime)
}

// 投票
func addArticleHasVoted(_r *redis.Client, _id string, _userID int64) bool {
	key := "voted:" + _id
	cmd := _r.SAdd(key, _userID)
	return cmd.Val() == 1
}

func checkArticleHasVoted(_r *redis.Client, _id string, _userID int64) bool {
	key := "voted:" + _id
	cmd := _r.SIsMember(key, _userID)
	return cmd.Val()
}

// 群组
func addGroups(_r *redis.Client, _id string, _groups []string) {
	for _, v := range _groups {
		key := fmt.Sprintf("group:%v", v)
		_r.SAdd(key, _id)
	}
}

func remGroups(_r *redis.Client, _id string, _groups []string) {
	for _, v := range _groups {
		key := fmt.Sprintf("group:%v", v)
		_r.SRem(key, _id)
	}
}

func getGroupArticle(_r *redis.Client, _groupID int64, _page int64) {
	key := fmt.Sprintf("order:group:%v", _groupID)
	if cmd := _r.Exists(key); cmd.Val() == 0 {
		_r.ZInterStore(key, redis.ZStore{}, fmt.Sprintf("group:%v", _groupID), "articlebyscore")
	}
	_r.Expire(key, 60*time.Second)

	ids := getArticleIDsTop(_r, key, _page)
	GetArticleByIDs(_r, ids)
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
