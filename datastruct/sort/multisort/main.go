package main

import (
	"fmt"
	"sort"
)

// 按订单金额由小到大排序，金额相同的按时间从大到小排序

type Order struct {
	Amount int64
	Time   int64
}

type Orders struct {
	Orders []*Order
}

func (orders *Orders) Len() int {
	return len(orders.Orders)
}

func (orders *Orders) Swap(i, j int) {
	orders.Orders[i], orders.Orders[j] = orders.Orders[j], orders.Orders[i]
}

func (orders *Orders) Less(i, j int) bool { // 这是一种完全错误的写法  什么都排序不出来
	if orders.Orders[i].Time > orders.Orders[j].Time {
		return true
	}
	if orders.Orders[i].Amount < orders.Orders[j].Amount {
		return true
	}
	return false
}

var myOrders *Orders

func init() {
	myOrders = &Orders{Orders: make([]*Order, 7)}
	myOrders.Orders[0] = &Order{
		Amount: 10,
		Time:   10,
	}
	myOrders.Orders[1] = &Order{
		Amount: 30,
		Time:   2,
	}
	myOrders.Orders[2] = &Order{
		Amount: 20,
		Time:   50,
	}
	myOrders.Orders[3] = &Order{
		Amount: 10,
		Time:   1,
	}
	myOrders.Orders[4] = &Order{
		Amount: 30,
		Time:   15,
	}
	myOrders.Orders[5] = &Order{
		Amount: 10,
		Time:   30,
	}
	myOrders.Orders[6] = &Order{
		Amount: 10,
		Time:   20,
	}
}

func main() {
	// sort.Sort(myOrders)
	sort.Slice(myOrders.Orders, func(i, j int) bool {
		return myOrders.Orders[i].Time > myOrders.Orders[j].Time
	})
	sort.Slice(myOrders.Orders, func(i, j int) bool {
		return myOrders.Orders[i].Amount < myOrders.Orders[j].Amount
	})
	for _, v := range myOrders.Orders {
		fmt.Println(v)
	}

}
