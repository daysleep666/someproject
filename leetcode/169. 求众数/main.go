package main

import (
	"fmt"
)

//摩尔投票法
func majorityElement(nums []int) int {
	var num = nums[0]
	var score = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			score++
		} else {
			score--
		}
		if score == 0 {
			num = nums[i]
			score = 1
		}
	}
	return num
}

func main() {
	fmt.Println(majorityElement([]int{4, 5, 4, 4, 4, 4, 3, 2, 1}))
}
