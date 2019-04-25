package main

import "fmt"

func findMaxAverage(nums []int, k int) float64 {
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	for first, last, cur := 0, k, max; last < len(nums); first, last = first+1, last+1 {
		if cur = cur - nums[first] + nums[last]; cur > max {
			max = cur
		}
	}
	return float64(max) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
}
