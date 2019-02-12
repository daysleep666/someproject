package distributionsort

import "fmt"

// 按订单金额由小到大排序，金额相同的按时间从大到小排序

type Order struct {
	Amount int64
	Time   int64
}

var myOrders []*Order

func init() {
	myOrders = make([]*Order, 10)
	myOrders[0] = &Order{
		Amount: 10,
		Time:   1549940110,
	}
	myOrders[1] = &Order{
		Amount: 30,
		Time:   1549940132,
	}
	myOrders[2] = &Order{
		Amount: 20,
		Time:   1549940150,
	}
	myOrders[3] = &Order{
		Amount: 10,
		Time:   1549940121,
	}
	myOrders[4] = &Order{
		Amount: 30,
		Time:   2549940115,
	}
	myOrders[5] = &Order{
		Amount: 10,
		Time:   1549940130,
	}
	myOrders[6] = &Order{
		Amount: 10,
		Time:   1549940120,
	}
	myOrders[7] = &Order{
		Amount: 30,
		Time:   1549941115,
	}
	myOrders[8] = &Order{
		Amount: 20,
		Time:   1549940330,
	}
	myOrders[9] = &Order{
		Amount: 10,
		Time:   1549940330,
	}
}

func TestRun() {
	distributionSortOrder(myOrders)
	for _, v := range myOrders {
		fmt.Println(v)
	}
}

func distributionSortOrder(_orders []*Order) {
	distributionSortByTime(_orders)
	distributionSortByAmount(_orders)
}

func distributionSortByTime(_orders []*Order) {
	// 时间戳
	var (
		m int64 = 10e10
	)
	for i := int64(1); i <= m; i *= 10 {
		bubble := make([][]*Order, 10)
		for _, v := range _orders {
			tmpV := v.Time / i
			tmpV = tmpV % 10
			bubble[tmpV] = append(bubble[tmpV], v)
		}

		var k int64
		for _, arr := range bubble {
			for _, v := range arr {
				_orders[k] = v
				k++
			}
		}
	}
}

func distributionSortByAmount(_orders []*Order) {
	var (
		max int64
	)

	for _, v := range _orders {
		if v.Amount > max {
			max = v.Amount
		}
	}

	for i := int64(1); i <= max; i *= 10 {
		bubble := make([][]*Order, 10)
		for _, v := range _orders {
			tmpV := v.Amount / i
			tmpV = tmpV % 10
			bubble[tmpV] = append(bubble[tmpV], v)
		}
		var k int64
		for _, arr := range bubble { //
			for _, v := range arr {
				_orders[k] = v
				k++
			}
		}
	}
}
