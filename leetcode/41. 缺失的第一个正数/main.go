package main

import "fmt"

func firstMissingPositive(nums []int) int {
	// arr := make([]int, len(nums))
	for _, v := range nums {
		if v >= 1 && v <= len(nums) {
			nums[v-1]++
		}
	}
	for i, v := range nums {
		if v == 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, 1}))
	fmt.Println(firstMissingPositive([]int{3, 4, 5}))
}
