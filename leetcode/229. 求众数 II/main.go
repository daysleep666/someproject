package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	count := len(nums) / 3
	// 超过count的就是众数
	sort.Ints(nums)
	result := []int{}
	cur := nums[0]
	num := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			num++
		} else {
			if num > count {
				result = append(result, cur)
			}
			cur = nums[i]
			num = 1
		}
	}
	if num > count {
		result = append(result, cur)
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
