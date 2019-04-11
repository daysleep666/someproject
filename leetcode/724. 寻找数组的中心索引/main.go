package main

import "fmt"

func pivotIndex(nums []int) int {
	arr := make([]int, len(nums)+1)
	sum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		sum += nums[i]
		arr[i] = sum
	}

	sum = 0
	for i := 0; i < len(nums); i++ {
		if sum == arr[i+1] {
			return i
		}
		sum += nums[i]
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))      // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))               // -1
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0})) // 2
	fmt.Println(pivotIndex([]int{}))                      // -1
	fmt.Println(pivotIndex([]int{-1, -1, -1, 0, 1, 1}))   // 0
	fmt.Println(pivotIndex([]int{-1, -1, 1, 1, 1}))       // 0
}
