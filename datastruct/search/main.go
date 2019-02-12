package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/search/binarysearch"
)

var arr = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	isFind := binarysearch.BinarySearch(arr, -9)
	fmt.Println(isFind)
}
