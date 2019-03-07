package main

import "fmt"

// 判断两个数相等用异或

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
}
