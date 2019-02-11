package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/sort/quicksort"
)

var arr []int64 = []int64{3, 2, 4, 1, 5, 8, 6, 0}

func main() {
	quicksort.QuickSort(arr)
	fmt.Println(arr)
}
