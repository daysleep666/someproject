package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var max, cur, value int = 1, 1, nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > value {
			cur++
		} else {
			cur = 1
		}
		value = nums[i]
		if cur > max {
			max = cur
		}
	}
	return max
}

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2}))
}
