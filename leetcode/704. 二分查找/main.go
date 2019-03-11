package main

import "fmt"

func search(nums []int, target int) int {
	from, to, mid := 0, len(nums)-1, 0
	for from <= to {
		mid = (from + to) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{5}, 5))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 8))
}
