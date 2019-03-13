package main

import (
	"fmt"
	"sort"
)

var result [][]int

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result = make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		k := len(nums)
		for j := 1; j < k; j++ {
			if v := bs(nums, j, k-1, -(nums[i] + nums[j])); v != -1 && !find(result, nums[i], nums[j], nums[v]) && i != j && i != v && j != v {
				result = append(result, []int{nums[i], nums[j], nums[v]})
				k = v
			}
		}
	}
	return result
}

func find(r [][]int, a, b, c int) bool {
	for _, v := range r {
		if (v[0] == a && v[1] == b) || (v[0] == a && v[1] == c) || (v[0] == b && v[1] == a) || (v[0] == b && v[1] == c) || (v[0] == c && v[1] == a) || (v[0] == c && v[1] == b) {
			return true
		}
	}
	return false
}

func bs(nums []int, low, high int, target int) int {
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if high >= 0 && nums[high] == target {
		return high
	}
	return -1
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-1, 0, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))

}
