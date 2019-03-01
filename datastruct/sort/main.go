package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/daysleep666/someproject/datastruct/sort/quicksort"
)

var arr []int64 = []int64{33, 211, 42, 12, 55, 71230, 453216, 21310}

func main() {
	daipai := make([]int64, 1000000)
	rand.Seed(11111)
	for idx := 0; idx != 1000000; idx++ {
		daipai[idx] = int64(rand.Int() % 200)
	}

	st := time.Now().UnixNano()
	quicksort.QuickSort(daipai)
	fmt.Println((time.Now().UnixNano() - st) / 1e6)
}
