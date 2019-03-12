package main

import "fmt"

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if f(nums, mid, low) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	qiyidian := low
	if target >= nums[qiyidian] && target <= nums[len(nums)-1] {
		searchh(nums, qiyidian, len(nums)-1, target)
	}
	return searchh(nums, 0, qiyidian-1, target)
}

func searchh(nums []int, low, high, target int) int {
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = high - 1
		} else {
			low = low + 1
		}
	}
	return -1
}

func f(nums []int, a, b int) bool {
	if nums[a] > b {
		return true
	}
	return false
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1))
}
