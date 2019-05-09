package main

import (
	"encoding/json"
	"fmt"

	"github.com/liunian1004/pdd"
)

func main() {
	{
		a := "aaa"
		defer fmt.Println(a, "222")
	}
	fmt.Println("111")

	return
	p := pdd.NewPdd(&pdd.Config{
		ClientId:     "",
		ClientSecret: "",
		RetryTimes:   1, // 设置接口调用失败重试次数
	})

	// 初始化多多客相关 API 调用
	d := p.GetDDK()

	// // 获取主题列表
	// r, err := d.ThemeListGet(1, 20)

	// 初始化商品 API
	// g := p.GetGoodsAPI()
	// for i := 1; i < 100; i++ {
	// 	gs, _ := d.GoodsSearch(pdd.Params{"page": i})
	// 	for _, v := range gs.GoodsList {
	// 		// r, _ := json.Marshal(v)
	// 		// fmt.Println(string(r))
	// 		if v.MinNormalPrice < 300 {
	// 			fmt.Println(v.GoodsName, v.GoodsDesc, v.MinNormalPrice, v.GoodsId)
	// 		}
	// 	}
	// }
	// 3864561472
	res, err := d.GoodsPromotionUrlGenerate("8472748_56301691", 8395580848, pdd.Params{"custom_parameters": "测试测试", "generate_weapp_webview": true, "generate_we_app": true})
	st, _ := json.Marshal(res)
	fmt.Println(res.Url, err)
	fmt.Println(string(st))
	// res1, _ := d.OrderListIncrementGet(time.Now().Unix()-86400, time.Now().Unix())
	// fmt.Println("count:", res1.TotalCount)
	// for _, v := range res1.OrderList {
	// 	fmt.Println(v.)
	// }
	// res2, _ := d.ThemeListGet(1, 10)
	// for _, v := range res2.ThemeList {
	// 	fmt.Println(v)
	// }

}

func test(a ...interface{}) {
	fmt.Println(len(a))
}
