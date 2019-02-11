package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/daysleep666/someproject/datastruct/sort/bubblesort"
	"github.com/daysleep666/someproject/datastruct/sort/insertionsort"
	"github.com/daysleep666/someproject/datastruct/sort/mergesort"
)

// 随机生成10000个数组每组200个数据

var bigArr [][]int64

var (
	big   = 100
	small = 5000
)

func init() {
	bigArr = make([][]int64, big)
	for i := 0; i < big; i++ {
		var arr []int64 = make([]int64, small)
		for j := 0; j < small; j++ {
			arr[j] = rand.Int63n(int64(small * 10))
		}
		bigArr[i] = arr
	}
}

// 冒泡排序 插入排序 选择排序

func main() {
	var st int64

	st = getCurMS()
	for _, v := range bigArr {
		bubblesort.BubbleSort(v)
	}
	fmt.Printf("冒泡排序:%vms\n", getCurMS()-st) // 162ms
	//------------------------------------------------------
	st = getCurMS()
	for _, v := range bigArr {
		insertionsort.InsertionSort(v)
	}
	fmt.Printf("插入排序:%vms\n", getCurMS()-st) // 18ms
	//------------------------------------------------------
	st = getCurMS()
	for _, v := range bigArr {
		mergesort.MergeSort(v)
	}
	fmt.Printf("归并排序:%vms\n", getCurMS()-st) // 8ms
	//------------------------------------------------------

}

func getCurMS() int64 {
	return time.Now().UnixNano() / 1e6
}
