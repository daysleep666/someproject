package main

import "fmt"

func searchRange(nums []int, target int) []int {
	ln := len(nums)
	if ln == 0 {
		return []int{-1, -1}
	}
	from, to, mid := 0, ln-1, 0
	for from < to {
		mid = (from+to)/2 + 1
		if nums[mid] == target {
			from++
		} else if nums[mid] > target {
			from = mid + 1
		} else {
			to = mid - 1
		}
	}
	if from >= ln || nums[to] != target {
		return []int{-1, -1}
	}

	min := from
	for i := from - 1; i >= 0; i-- {
		if nums[i] == target {
			min = i
		} else {
			break
		}
	}
	return []int{min, from}
}

func main() {
	fmt.Println(searchRange([]int{1, 4}, 4))
	fmt.Println(searchRange([]int{1}, 1))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{2, 2}, 6))
	fmt.Println(searchRange([]int{2, 2}, 1))
	fmt.Println(searchRange([]int{1, 3}, 1))
}
