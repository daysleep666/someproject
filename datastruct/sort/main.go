package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/sort/insertionsort"
)

var arr []int64 = []int64{3, 2, 4, 1, 5, 8, 6, 0}

func main() {
	arr = insertionsort.InsertionSort(arr)
	fmt.Println(arr)
}
