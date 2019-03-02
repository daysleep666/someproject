package main

import "fmt"

func productExceptSelf(nums []int) []int { // 时间复杂度O(n) 空间复杂度O(1)
	lenNum := len(nums)
	output := make([]int, lenNum)
	left := 1
	for i, v := range nums {
		output[i] = left
		left *= v
	}
	right := 1
	for i := lenNum - 1; i >= 0; i-- {
		output[i] *= right
		right *= nums[i]
	}
	return output
}

func main() {
	num := []int{1, 2, 3}
	fmt.Println(productExceptSelf(num))
}
