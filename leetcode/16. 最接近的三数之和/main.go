package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if cur, min := compute(nums[i]+nums[j]+nums[k], target), compute(result, target); cur < min {
					result = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return result
}

func compute(a, target int) int {
	if a-target > 0 {
		return a - target
	}
	return target - a
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
	fmt.Println(threeSumClosest([]int{0, 5, -1, -2, 4, -1, 0, -3, 4, -5}, 1))
}
