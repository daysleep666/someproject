package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/smartwalle/alipay"
)

var prikey = `MIIEpAIBAAKCAQEAp7EkbH6NORAUjEiHBIddm0p6CcJw9sk0acLNnUs30g9t0b7l2DkPrDTO0m/D/gf/DGywssW6PBRM7KxfT7vr15Ga2lWb1axD2+DJF8i7euBbbEfUXoDW2zHI6HAMUnntQUgkzUNI5zPjHyAwqguVb52pLuY7R4l0Rqa6VZHn22roX9u2BzpGzQU9GdaXnYX9W86Is9Nq+t4Ht04bRmBqZHjNE4ppjCqyC8QyjG85izmqu5cUUNAwzBNXovmf/zdeQfBJ4l+qunv/4OKfhM0AtUM2xQikZQsri4vzyZPdOSVmMyXNeuiKSahvltRFgvw5Fddhk61MWb2Go2bDrWYXkwIDAQABAoIBADH4NltiJqEPaqAW6KzvzXP98HnCCDlx4X12gCDUDjE5ScdLtGAHOOVe4xaQGX40UcsZBFpxUzqd7gJEpy4yWGfnSpavFf3JFHuL80B0H3HRM00D5oO2rs+oeS7SXMmi7ZEbd2P3Hd1Pd+mJhRbW1tPozkZccuOxUQXQnwvqq4WVTVnkGZPYAmHxAH13EXVpQgytdb/ohzaxQ0Y4PzObyCxcFEks1c8vzVoI9N8vO4iOrSliVsNCdO6nflQ4XiDJNj2hNS0Oqz+cJCbEjZHvJPa/iy+s12YPSm712+PYYWyUAVRT75ME1TUi/KBOhbJWENSIv2cFoS5yPE2e3QdNTrECgYEA1XEGLI9a9Mu3uNgztshTuger+Nt+A//Sh/36PLRpmnMM0hDO0lR84u7mNhLoY6mdR3nihvbX1w7q0Hiek3VD6odrZFrNzaPgphrMojZB7CYgumfDPVRSmrGAReeDFTHSw+wbzUqSOFRRfv6sv98e6L5cgAKljsWlt6HYon2u7isCgYEAySDdOcSdvAdrZhmpFGVsDaLfUlxBin0/W3+M+oG4/pT9LV76KKUuW+28vmczCQINXL2nmUx6dw/G5QBYoeXCiwm3ULPx3ijZW6ZGYeWL14obf+1ShgDPoXUKB4zvHVqOtlsNdnTJGNgXbmURWuuws+T+evReHHVVofCx+TPhMDkCgYEAkXz5OGZ/GV3DXbwBhBYI/EU+c+MpgpBecVQtvjM7yIujsYzDbZsphYZSZYl93nE/Q/lQ765SccU5SeDuNpU4NiUYjs9WeVW7S334/pj9Vnqm1GjmLGfN5qQXS48DSu8rO2E1kjnCcaMH6giuuBredsdLRfaH40nWK9WmVRxWoQECgYEAlZa60ydYBawjBdoZ9hCE8BNuM4xoNr92qu/QA9UNqVNYhjE4r5hDRZ4pHhgVTiLXm113VL4b64exb0r1oNwI+G4uiHh9cLVDDDY77L3UZ2fRkjEJip75zmUFI55/jX533EsWVT8A/lhh/PpwkQXw86tiVkJ/6HygRwSE6gra9okCgYB2PRrQxoBJdyp9dmNws3wSavglf/nFiTKql6nMuPOx2a+i4rlfiVZBfCbAFti2AG7qPl/WlE8qfXk4nBesHTGlIGICF7cjswa5Rxl+jntfNQggSfyhHFZMXMLerzd9k3PWhpn7VRqrfIAGEaMY4WqsmBNw2CktDdJbLFM49Di3OQ==`

var pubkey = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp7EkbH6NORAUjEiHBIddm0p6CcJw9sk0acLNnUs30g9t0b7l2DkPrDTO0m/D/gf/DGywssW6PBRM7KxfT7vr15Ga2lWb1axD2+DJF8i7euBbbEfUXoDW2zHI6HAMUnntQUgkzUNI5zPjHyAwqguVb52pLuY7R4l0Rqa6VZHn22roX9u2BzpGzQU9GdaXnYX9W86Is9Nq+t4Ht04bRmBqZHjNE4ppjCqyC8QyjG85izmqu5cUUNAwzBNXovmf/zdeQfBJ4l+qunv/4OKfhM0AtUM2xQikZQsri4vzyZPdOSVmMyXNeuiKSahvltRFgvw5Fddhk61MWb2Go2bDrWYXkwIDAQAB`

func main() {
	var appId = "2016092200572282"
	var aliPublicKey = ""   // 可选，支付宝提供给我们用于签名验证的公钥，通过支付宝管理后台获取
	var privateKey = prikey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client = alipay.New(appId, aliPublicKey, privateKey, false)

	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://api.wancai188.cn/callback/ali"
	// p.ReturnURL = "http://gogs.wancai188.cn:1234/callback/ali"
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
