package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// a()
	// b()
	pipelineTest(getClient())
	time.Sleep(time.Second * 3)
}

func a() {
	client := getClient()
	addItem(client, 1, "car")
	AddToMarket(client, 1, "car", 100)
	incFunds(client, 2, 100)
	incFunds(client, 3, 100)
	incFunds(client, 4, 100)
	incFunds(client, 5, 100)
}

func b() {
	for i := int64(2); i <= 5; i++ { // 预测只有一个success
		go func(index int64) {
			client := getClient()
			PurchaseItem(client, index, 1, "car")
		}(i)
	}
}

func pipelineTest(r *redis.Client) {
	r.Del("a", "b")
	r.Set("a", 0, time.Second)
	err := r.Watch(func(tx *redis.Tx) error {
		p := tx.TxPipeline()
		p.Set("b", "a", 0)
		time.Sleep(time.Second * 2)
		_, err := p.Exec()
		return err
	}, "a")
	fmt.Println("result =", err)
	// p := r.Pipeline()
	// p.Set("a", "a", 0)
	// fmt.Println("a:", p.Get("a").Val())
	// p.Set("b", "b", 0)
	// p.Exec()
	// p := r.Pipeline()
	// p.Set("hello", "world", 0)
	// tp := p.TxPipeline()
	// tp.Set("hey", "qwe", 0)
	// tp.Exec()
}

// pipleline里没发查询key
func AddToMarket(r *redis.Client, uID int64, itemID string, price int64) {
	end := time.Now().Unix() + 2
	inverntoryKey := fmt.Sprintf("inventory:%v", uID)
	marketKey := "market"
	for time.Now().Unix() < end { // 在五秒内连续尝试直到成功
		err := r.Watch(func(tx *redis.Tx) error {
			// time.Sleep(time.Second)
			// 在此时修改inverntoryKey值，会导致事务失败[redis: transaction failed]

			//
			if cmd := tx.SIsMember(inverntoryKey, itemID); !cmd.Val() { // 如果这里存在这个物品
				return fmt.Errorf("%v hav't %v in %v (%v)", uID, itemID, inverntoryKey, cmd.Val())
			}

			// 开启事务
			txp := tx.TxPipeline()
			txp.ZAdd(marketKey, redis.Z{Member: fmt.Sprintf("%v:%v", itemID, uID), Score: float64(price)})
			txp.SRem(inverntoryKey, itemID)
			_, err := txp.Exec()
			return err
		}, inverntoryKey)
		if err != nil { // 说明成功了
			fmt.Printf("Result:%v\n", err.Error())
		} else {
			fmt.Println("success")
			break
		}
	}

}

//
func PurchaseItem(r *redis.Client, buyerUID, sellerUID int64, itemID string) {
	end := time.Now().Unix() + 1
	buyerKey := fmt.Sprintf("user:%v", buyerUID)
	sellerKey := fmt.Sprintf("user:%v", sellerUID)
	buyerItemKey := fmt.Sprintf("inventory:%v", buyerUID)
	marketKey := "market"
	for time.Now().Unix() < end {
		err := r.Watch(func(tx *redis.Tx) error {
			// 获得商品价格
			price := int64(tx.ZScore(marketKey, fmt.Sprintf("%v:%v", itemID, sellerUID)).Val())
			// 检查买家钱够不够
			cmd := tx.HGet(buyerKey, "funds")
			funds, _ := cmd.Int64()
			if price == 0 {
				return fmt.Errorf("not exist %v", itemID)
			}
			if funds < price {
				return fmt.Errorf("not enough money %v %v", funds, price)
			}
			// 开启事务
			txp := tx.TxPipeline()
			// 转移钱
			txp.HIncrBy(buyerKey, "funds", -price)
			txp.HIncrBy(sellerKey, "funds", price)
			// 转移物品
			txp.SAdd(buyerItemKey, itemID)
			txp.ZRem(marketKey, fmt.Sprintf("%v:%v", itemID, sellerUID))
			_, err := txp.Exec()
			return err
		},
			fmt.Sprintf("user:%v", buyerUID),
			marketKey)
		if err != nil { // 说明成功了
			// fmt.Printf("Result:%v\n", err.Error())
		} else {
			fmt.Printf("祝贺 %v 购买成功了\n", buyerUID)
			break
		}
	}
}

func addUser(r *redis.Client, uID int64, name string) {
	// hashMap
	key := fmt.Sprintf("user:%v", uID)
	r.HSet(key, "name", name)
}

func incFunds(r *redis.Client, uID int64, funds int64) {
	key := fmt.Sprintf("user:%v", uID)
	r.HIncrBy(key, "funds", funds)
}

func addItem(r *redis.Client, uID int64, itemID string) {
	// set
	key := fmt.Sprintf("inventory:%v", uID)
	r.SAdd(key, itemID)
}

func existItem(r *redis.Client, uID int64, itemID string) bool {
	key := fmt.Sprintf("inventory:%v", uID)
	cmd := r.SIsMember(key, itemID)
	return cmd.Val()
}

func generateUserID(r *redis.Client) int64 {
	cmd := r.Incr("userid")
	return cmd.Val()
}

func addToMarket(r *redis.Client, uID int64, itemID string, price int64) {
	key := "market"
	r.ZAdd(key, redis.Z{Member: fmt.Sprintf("%v:%v", itemID, uID), Score: float64(price)})
}

//------------------------------------------------

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
