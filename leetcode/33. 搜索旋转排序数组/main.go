package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if f(nums, mid, 0) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	qiyidian := low
	fmt.Println("qiyidian =", qiyidian)
	if qiyidian >= 0 && qiyidian <= len(nums)-1 && target >= nums[qiyidian] && target <= nums[len(nums)-1] {
		return searchh(nums, qiyidian, len(nums)-1, target)
	}
	return searchh(nums, 0, qiyidian-1, target)
}

func searchh(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func f(nums []int, a, b int) bool {
	if nums[a] < nums[b] {
		return false
	}
	return true
}

func main() {
	fmt.Println(search([]int{5, 1, 3}, 1))             // 1
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 1)) // 4
	fmt.Println(search([]int{4, 6}, 6))                // 2
	fmt.Println(search([]int{4, 6, 1}, 6))             // 2
}
