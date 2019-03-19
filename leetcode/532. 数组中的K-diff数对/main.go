package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	pairs := 0
	sort.Ints(nums)
	for i, j := 0, 1; i < len(nums)-1 && j < len(nums); {
		if nums[j]-nums[i] == k {
			i++
			j++
			for i < len(nums)-1 && j < len(nums) && nums[i] == nums[i-1] && nums[j] == nums[j-1] {
				i++
				j++
			}
			pairs++
		} else if nums[j]-nums[i] > k {
			i++
			for i >= j {
				j++
			}
		} else {
			j++
		}
	}
	return pairs
}

func main() {
	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}
