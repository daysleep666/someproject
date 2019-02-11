package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/daysleep666/someproject/datastruct/sort/bubblesort"
	"github.com/daysleep666/someproject/datastruct/sort/insertionsort"
)

// 随机生成10000个数组每组200个数据

var bigArr [][]int64

func init() {
	bigArr = make([][]int64, 10000)
	for i := 0; i < 10000; i++ {
		var arr []int64 = make([]int64, 200)
		for j := 0; j < 200; j++ {
			arr[j] = rand.Int63n(300)
		}
		bigArr[i] = arr
	}
}

// 冒泡排序 插入排序 选择排序

func main() {
	st := getCurMS()
	for _, v := range bigArr {
		bubblesort.BubbleSort(v)
	}
	fmt.Printf("冒泡排序:%v\n", getCurMS()-st)

	st = getCurMS()
	for _, v := range bigArr {
		insertionsort.InsertionSort(v)
	}
	fmt.Printf("插入排序:%v\n", getCurMS()-st)
}

func getCurMS() int64 {
	return time.Now().UnixNano() / 1e6
}
