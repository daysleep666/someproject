package main

import (
	"fmt"
)

func FindMin(nums []int, target int) int { // 找到第一个等于targe的值
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] >= target {
			high = high - 1
		} else {
			low = low + 1
		}
	}

	if low >= len(nums) || nums[low] != target {
		return -1
	}
	return low
}

func FindMax(nums []int, target int) int { // 找到最后一个等于target的值
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if high < 0 || nums[high] != target {
		return -1
	}
	return high
}

func main() {
	fmt.Println(FindMax([]int{1, 2, 3, 4}, 2))
	fmt.Println(FindMax([]int{1, 2, 2, 2, 3, 4}, 2))
	fmt.Println(FindMax([]int{1, 2, 2, 2, 3, 4}, 0))
	fmt.Println(FindMax([]int{1, 2, 2, 2, 3, 4}, 5))
}
