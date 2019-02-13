package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/search/binarysearch"
)

var arr = []int64{0, 1, 2, 3}

func main() {
	isFind := binarysearch.BinarySearch(arr, 3)
	fmt.Println(isFind)
}
