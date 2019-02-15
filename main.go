package main

import (
	"fmt"
)

func main() {
	a := testSlice()
	fmt.Printf("切片函数外%p\n", &a)

	b := testArray()
	fmt.Printf("数组函数外%p\n", &b)
}

func testSlice() []int64 {
	var a = make([]int64, 0)
	fmt.Printf("切片函数里%p\n", &a)
	return a
}

func testArray() [3]int64 {
	var a [3]int64
	fmt.Printf("数组函数里%p\n", &a)
	return a
}
