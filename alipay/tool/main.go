package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/smartwalle/alipay"
)

var prikey = ``

var pubkey = ``

func main() {
	var appId = "00000"
	var aliPublicKey = ""   // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
	var privateKey = prikey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client = alipay.New(appId, aliPublicKey, privateKey, true)

	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = ""
	// p.ReturnURL = ""
	p.Subject = "标题"
	p.OutTradeNo = bson.NewObjectId().Hex()
	p.TotalAmount = "1.00"
	p.ProductCode = "QUICK_WAP_WAY"
	p.TimeoutExpress = "1c"
	var url, err = client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}
	var payURL = url.String()
	fmt.Println(payURL)
	// 这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
}
