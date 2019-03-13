package main

import (
	"fmt"
)

func LE(data []int, a, b int) bool {
	return data[a] <= data[b]
}

func CheckMin(data []int, lo, hi int) bool {
	if lo >= hi {
		return true
	}

	for cur := lo; cur <= (hi-lo-1)/2; cur++ {
		l := (cur-lo)*2 + 1
		if !LE(data, cur, l) {
			return false
		}
		r := l + 1
		if r > hi {
			break
		}
		if !LE(data, cur, r) {
			return false
		}
	}
	return true
}

// 大顶堆
func heapSort(nums []int) {
	heapify(nums, 0, len(nums)-1)
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, 0, i-1)
	}
}

func heapify(nums []int, low, high int) {
	for i := (high - 1) / 2; i >= low; i-- {
		down(nums, i, high)
	}
}

func down(nums []int, low, high int) {
	i := low
	for {
		childIndex := i
		if i*2+1 <= high && nums[i] < nums[i*2+1] {
			childIndex = i*2 + 1
		}
		if i*2+2 <= high && nums[childIndex] < nums[i*2+2] {
			childIndex = i*2 + 2
		}
		if childIndex == i {
			break
		}
		nums[i], nums[childIndex] = nums[childIndex], nums[i]
	}
}

func main() {
	// n := 5
	// arr := make([]int, 0)
	// for i := n; i >= 0; i-- {
	// 	arr = append(arr, rand.Intn(10))
	// }
	arr := []int{5, 4, 3, 2, 1, 10}
	// heapify(arr, 0, len(arr)-1)
	heapSort(arr)
	fmt.Println(arr)
}

// func isPowerOfTwo(x uintptr) bool {
// 	return x&(x-1) == 0
// }
